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

	"github.com/livekit/gstplugins/base"
)

type Rpicamsrc struct {
	*base.GstPushSrc
}

func NewRpicamsrc() (*Rpicamsrc, error) {
	e, err := gst.NewElement("rpicamsrc")
	if err != nil {
		return nil, err
	}
	return &Rpicamsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewRpicamsrcWithName(name string) (*Rpicamsrc, error) {
	e, err := gst.NewElementWithName("rpicamsrc", name)
	if err != nil {
		return nil, err
	}
	return &Rpicamsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// annotation-mode (GstRpiCamSrcAnnotationMode)
//
// Flags to control annotation of the output video

func (e *Rpicamsrc) GetAnnotationMode() (GstRpiCamSrcAnnotationMode, error) {
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

func (e *Rpicamsrc) SetAnnotationMode(value GstRpiCamSrcAnnotationMode) error {
	e.Element.SetArg("annotation-mode", string(value))
	return nil
}

// annotation-text (string)
//
// Text string to annotate onto video when annotation-mode flags include 'custom-text'

func (e *Rpicamsrc) GetAnnotationText() (string, error) {
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

func (e *Rpicamsrc) SetAnnotationText(value string) error {
	return e.Element.SetProperty("annotation-text", value)
}

// annotation-text-bg-colour (int)
//
// Set the annotation text background colour, as the integer corresponding to a VUY value eg 0x8080FF = 8421631, -1 for default

func (e *Rpicamsrc) GetAnnotationTextBgColour() (int, error) {
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

func (e *Rpicamsrc) SetAnnotationTextBgColour(value int) error {
	return e.Element.SetProperty("annotation-text-bg-colour", value)
}

// annotation-text-colour (int)
//
// Set the annotation text colour, as the integer corresponding to a VUY value eg 0x8080FF = 8421631, -1 for default

func (e *Rpicamsrc) GetAnnotationTextColour() (int, error) {
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

func (e *Rpicamsrc) SetAnnotationTextColour(value int) error {
	return e.Element.SetProperty("annotation-text-colour", value)
}

// annotation-text-size (int)
//
// Set the size of annotation text (in pixels) (0 = Auto)

func (e *Rpicamsrc) GetAnnotationTextSize() (int, error) {
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

func (e *Rpicamsrc) SetAnnotationTextSize(value int) error {
	return e.Element.SetProperty("annotation-text-size", value)
}

// awb-gain-blue (float32)
//
// Manual AWB Gain for blue channel when awb-mode=off

func (e *Rpicamsrc) GetAwbGainBlue() (float32, error) {
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

func (e *Rpicamsrc) SetAwbGainBlue(value float32) error {
	return e.Element.SetProperty("awb-gain-blue", value)
}

// awb-gain-red (float32)
//
// Manual AWB Gain for red channel when awb-mode=off

func (e *Rpicamsrc) GetAwbGainRed() (float32, error) {
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

func (e *Rpicamsrc) SetAwbGainRed(value float32) error {
	return e.Element.SetProperty("awb-gain-red", value)
}

// awb-mode (GstRpiCamSrcAWBMode)
//
// White Balance mode

func (e *Rpicamsrc) GetAwbMode() (interface{}, error) {
	return e.Element.GetProperty("awb-mode")
}

func (e *Rpicamsrc) SetAwbMode(value interface{}) error {
	return e.Element.SetProperty("awb-mode", value)
}

// bitrate (int)
//
// Bitrate for encoding. 0 for VBR using quantisation-parameter

func (e *Rpicamsrc) GetBitrate() (int, error) {
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

func (e *Rpicamsrc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// brightness (int)
//
// Image capture brightness

func (e *Rpicamsrc) GetBrightness() (int, error) {
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

func (e *Rpicamsrc) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// camera-number (int)
//
// Which camera to use on a multi-camera system - 0 or 1

func (e *Rpicamsrc) GetCameraNumber() (int, error) {
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

func (e *Rpicamsrc) SetCameraNumber(value int) error {
	return e.Element.SetProperty("camera-number", value)
}

// contrast (int)
//
// Image capture contrast

func (e *Rpicamsrc) GetContrast() (int, error) {
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

func (e *Rpicamsrc) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// drc (GstRpiCamSrcDRCLevel)
//
// Dynamic Range Control level

func (e *Rpicamsrc) GetDrc() (interface{}, error) {
	return e.Element.GetProperty("drc")
}

func (e *Rpicamsrc) SetDrc(value interface{}) error {
	return e.Element.SetProperty("drc", value)
}

// exposure-compensation (int)
//
// Exposure Value compensation

func (e *Rpicamsrc) GetExposureCompensation() (int, error) {
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

func (e *Rpicamsrc) SetExposureCompensation(value int) error {
	return e.Element.SetProperty("exposure-compensation", value)
}

// exposure-mode (GstRpiCamSrcExposureMode)
//
// Camera exposure mode to use

func (e *Rpicamsrc) GetExposureMode() (GstRpiCamSrcExposureMode, error) {
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

func (e *Rpicamsrc) SetExposureMode(value GstRpiCamSrcExposureMode) error {
	e.Element.SetArg("exposure-mode", string(value))
	return nil
}

// fullscreen (bool)
//
// Display preview window full screen

func (e *Rpicamsrc) GetFullscreen() (bool, error) {
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

func (e *Rpicamsrc) SetFullscreen(value bool) error {
	return e.Element.SetProperty("fullscreen", value)
}

// hflip (bool)
//
// Flip capture horizontally

func (e *Rpicamsrc) GetHflip() (bool, error) {
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

func (e *Rpicamsrc) SetHflip(value bool) error {
	return e.Element.SetProperty("hflip", value)
}

// image-effect (GstRpiCamSrcImageEffect)
//
// Visual FX to apply to the image

func (e *Rpicamsrc) GetImageEffect() (GstRpiCamSrcImageEffect, error) {
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

func (e *Rpicamsrc) SetImageEffect(value GstRpiCamSrcImageEffect) error {
	e.Element.SetArg("image-effect", string(value))
	return nil
}

// inline-headers (bool)
//
// Set to TRUE to insert SPS/PPS before each IDR packet

func (e *Rpicamsrc) GetInlineHeaders() (bool, error) {
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

func (e *Rpicamsrc) SetInlineHeaders(value bool) error {
	return e.Element.SetProperty("inline-headers", value)
}

// intra-refresh-type (GstRpiCamSrcIntraRefreshType)
//
// Type of Intra Refresh to use, -1 to disable intra refresh

func (e *Rpicamsrc) GetIntraRefreshType() (GstRpiCamSrcIntraRefreshType, error) {
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

func (e *Rpicamsrc) SetIntraRefreshType(value GstRpiCamSrcIntraRefreshType) error {
	e.Element.SetArg("intra-refresh-type", string(value))
	return nil
}

// iso (int)
//
// ISO value to use (0 = Auto)

func (e *Rpicamsrc) GetIso() (int, error) {
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

func (e *Rpicamsrc) SetIso(value int) error {
	return e.Element.SetProperty("iso", value)
}

// keyframe-interval (int)
//
// Interval (in frames) between I frames. -1 = automatic, 0 = single-keyframe

func (e *Rpicamsrc) GetKeyframeInterval() (int, error) {
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

func (e *Rpicamsrc) SetKeyframeInterval(value int) error {
	return e.Element.SetProperty("keyframe-interval", value)
}

// metering-mode (GstRpiCamSrcExposureMeteringMode)
//
// Camera exposure metering mode to use

func (e *Rpicamsrc) GetMeteringMode() (GstRpiCamSrcExposureMeteringMode, error) {
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

func (e *Rpicamsrc) SetMeteringMode(value GstRpiCamSrcExposureMeteringMode) error {
	e.Element.SetArg("metering-mode", string(value))
	return nil
}

// preview (bool)
//
// Display preview window overlay

func (e *Rpicamsrc) GetPreview() (bool, error) {
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

func (e *Rpicamsrc) SetPreview(value bool) error {
	return e.Element.SetProperty("preview", value)
}

// preview-encoded (bool)
//
// Display encoder output in the preview

func (e *Rpicamsrc) GetPreviewEncoded() (bool, error) {
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

func (e *Rpicamsrc) SetPreviewEncoded(value bool) error {
	return e.Element.SetProperty("preview-encoded", value)
}

// preview-h (int)
//
// Height of the preview window (in pixels)

func (e *Rpicamsrc) GetPreviewH() (int, error) {
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

func (e *Rpicamsrc) SetPreviewH(value int) error {
	return e.Element.SetProperty("preview-h", value)
}

// preview-opacity (int)
//
// Opacity to use for the preview window

func (e *Rpicamsrc) GetPreviewOpacity() (int, error) {
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

func (e *Rpicamsrc) SetPreviewOpacity(value int) error {
	return e.Element.SetProperty("preview-opacity", value)
}

// preview-w (int)
//
// Width of the preview window (in pixels)

func (e *Rpicamsrc) GetPreviewW() (int, error) {
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

func (e *Rpicamsrc) SetPreviewW(value int) error {
	return e.Element.SetProperty("preview-w", value)
}

// preview-x (int)
//
// Start X coordinate of the preview window (in pixels)

func (e *Rpicamsrc) GetPreviewX() (int, error) {
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

func (e *Rpicamsrc) SetPreviewX(value int) error {
	return e.Element.SetProperty("preview-x", value)
}

// preview-y (int)
//
// Start Y coordinate of the preview window (in pixels)

func (e *Rpicamsrc) GetPreviewY() (int, error) {
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

func (e *Rpicamsrc) SetPreviewY(value int) error {
	return e.Element.SetProperty("preview-y", value)
}

// quantisation-parameter (int)
//
// Set a Quantisation Parameter approx 10-40 with bitrate=0 for VBR encoding. 0 = off

func (e *Rpicamsrc) GetQuantisationParameter() (int, error) {
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

func (e *Rpicamsrc) SetQuantisationParameter(value int) error {
	return e.Element.SetProperty("quantisation-parameter", value)
}

// roi-h (float32)
//
// Normalised region-of-interest H coord

func (e *Rpicamsrc) GetRoiH() (float32, error) {
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

func (e *Rpicamsrc) SetRoiH(value float32) error {
	return e.Element.SetProperty("roi-h", value)
}

// roi-w (float32)
//
// Normalised region-of-interest W coord

func (e *Rpicamsrc) GetRoiW() (float32, error) {
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

func (e *Rpicamsrc) SetRoiW(value float32) error {
	return e.Element.SetProperty("roi-w", value)
}

// roi-x (float32)
//
// Normalised region-of-interest X coord

func (e *Rpicamsrc) GetRoiX() (float32, error) {
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

func (e *Rpicamsrc) SetRoiX(value float32) error {
	return e.Element.SetProperty("roi-x", value)
}

// roi-y (float32)
//
// Normalised region-of-interest Y coord

func (e *Rpicamsrc) GetRoiY() (float32, error) {
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

func (e *Rpicamsrc) SetRoiY(value float32) error {
	return e.Element.SetProperty("roi-y", value)
}

// rotation (int)
//
// Rotate captured image (0, 90, 180, 270 degrees)

func (e *Rpicamsrc) GetRotation() (int, error) {
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

func (e *Rpicamsrc) SetRotation(value int) error {
	return e.Element.SetProperty("rotation", value)
}

// saturation (int)
//
// Image capture saturation

func (e *Rpicamsrc) GetSaturation() (int, error) {
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

func (e *Rpicamsrc) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

// sensor-mode (GstRpiCamSrcSensorMode)
//
// Manually set the camera sensor mode

func (e *Rpicamsrc) GetSensorMode() (GstRpiCamSrcSensorMode, error) {
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

func (e *Rpicamsrc) SetSensorMode(value GstRpiCamSrcSensorMode) error {
	e.Element.SetArg("sensor-mode", string(value))
	return nil
}

// sharpness (int)
//
// Image capture sharpness

func (e *Rpicamsrc) GetSharpness() (int, error) {
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

func (e *Rpicamsrc) SetSharpness(value int) error {
	return e.Element.SetProperty("sharpness", value)
}

// shutter-speed (int)
//
// Set a fixed shutter speed, in microseconds. (0 = Auto)

func (e *Rpicamsrc) GetShutterSpeed() (int, error) {
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

func (e *Rpicamsrc) SetShutterSpeed(value int) error {
	return e.Element.SetProperty("shutter-speed", value)
}

// use-stc (bool)
//
// Use the camera STC for timestamping buffers

func (e *Rpicamsrc) GetUseStc() (bool, error) {
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

func (e *Rpicamsrc) SetUseStc(value bool) error {
	return e.Element.SetProperty("use-stc", value)
}

// vflip (bool)
//
// Flip capture vertically

func (e *Rpicamsrc) GetVflip() (bool, error) {
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

func (e *Rpicamsrc) SetVflip(value bool) error {
	return e.Element.SetProperty("vflip", value)
}

// video-stabilisation (bool)
//
// Enable or disable video stabilisation

func (e *Rpicamsrc) GetVideoStabilisation() (bool, error) {
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

func (e *Rpicamsrc) SetVideoStabilisation(value bool) error {
	return e.Element.SetProperty("video-stabilisation", value)
}

// ----- Constants -----

type GstRpiCamSrcAwbmode string

const (
	GstRpiCamSrcAwbmodeOff GstRpiCamSrcAwbmode = "off" // off (0) – GST_RPI_CAM_SRC_AWB_MODE_OFF
	GstRpiCamSrcAwbmodeAuto GstRpiCamSrcAwbmode = "auto" // auto (1) – GST_RPI_CAM_SRC_AWB_MODE_AUTO
	GstRpiCamSrcAwbmodeSunlight GstRpiCamSrcAwbmode = "sunlight" // sunlight (2) – GST_RPI_CAM_SRC_AWB_MODE_SUNLIGHT
	GstRpiCamSrcAwbmodeCloudy GstRpiCamSrcAwbmode = "cloudy" // cloudy (3) – GST_RPI_CAM_SRC_AWB_MODE_CLOUDY
	GstRpiCamSrcAwbmodeShade GstRpiCamSrcAwbmode = "shade" // shade (4) – GST_RPI_CAM_SRC_AWB_MODE_SHADE
	GstRpiCamSrcAwbmodeTungsten GstRpiCamSrcAwbmode = "tungsten" // tungsten (5) – GST_RPI_CAM_SRC_AWB_MODE_TUNGSTEN
	GstRpiCamSrcAwbmodeFluorescent GstRpiCamSrcAwbmode = "fluorescent" // fluorescent (6) – GST_RPI_CAM_SRC_AWB_MODE_FLUORESCENT
	GstRpiCamSrcAwbmodeIncandescent GstRpiCamSrcAwbmode = "incandescent" // incandescent (7) – GST_RPI_CAM_SRC_AWB_MODE_INCANDESCENT
	GstRpiCamSrcAwbmodeFlash GstRpiCamSrcAwbmode = "flash" // flash (8) – GST_RPI_CAM_SRC_AWB_MODE_FLASH
	GstRpiCamSrcAwbmodeHorizon GstRpiCamSrcAwbmode = "horizon" // horizon (9) – GST_RPI_CAM_SRC_AWB_MODE_HORIZON
)

type GstRpiCamSrcAnnotationMode string

const (
	GstRpiCamSrcAnnotationModeCustomText GstRpiCamSrcAnnotationMode = "custom-text" // custom-text (0x00000001) – GST_RPI_CAM_SRC_ANNOTATION_MODE_CUSTOM_TEXT
	GstRpiCamSrcAnnotationModeText GstRpiCamSrcAnnotationMode = "text" // text (0x00000002) – GST_RPI_CAM_SRC_ANNOTATION_MODE_TEXT
	GstRpiCamSrcAnnotationModeDate GstRpiCamSrcAnnotationMode = "date" // date (0x00000004) – GST_RPI_CAM_SRC_ANNOTATION_MODE_DATE
	GstRpiCamSrcAnnotationModeTime GstRpiCamSrcAnnotationMode = "time" // time (0x00000008) – GST_RPI_CAM_SRC_ANNOTATION_MODE_TIME
	GstRpiCamSrcAnnotationModeShutterSettings GstRpiCamSrcAnnotationMode = "shutter-settings" // shutter-settings (0x00000010) – GST_RPI_CAM_SRC_ANNOTATION_MODE_SHUTTER_SETTINGS
	GstRpiCamSrcAnnotationModeCafSettings GstRpiCamSrcAnnotationMode = "caf-settings" // caf-settings (0x00000020) – GST_RPI_CAM_SRC_ANNOTATION_MODE_CAF_SETTINGS
	GstRpiCamSrcAnnotationModeGainSettings GstRpiCamSrcAnnotationMode = "gain-settings" // gain-settings (0x00000040) – GST_RPI_CAM_SRC_ANNOTATION_MODE_GAIN_SETTINGS
	GstRpiCamSrcAnnotationModeLensSettings GstRpiCamSrcAnnotationMode = "lens-settings" // lens-settings (0x00000080) – GST_RPI_CAM_SRC_ANNOTATION_MODE_LENS_SETTINGS
	GstRpiCamSrcAnnotationModeMotionSettings GstRpiCamSrcAnnotationMode = "motion-settings" // motion-settings (0x00000100) – GST_RPI_CAM_SRC_ANNOTATION_MODE_MOTION_SETTINGS
	GstRpiCamSrcAnnotationModeFrameNumber GstRpiCamSrcAnnotationMode = "frame-number" // frame-number (0x00000200) – GST_RPI_CAM_SRC_ANNOTATION_MODE_FRAME_NUMBER
	GstRpiCamSrcAnnotationModeBlackBackground GstRpiCamSrcAnnotationMode = "black-background" // black-background (0x00000400) – GST_RPI_CAM_SRC_ANNOTATION_MODE_BLACK_BACKGROUND
)

type GstRpiCamSrcDrclevel string

const (
	GstRpiCamSrcDrclevelOff GstRpiCamSrcDrclevel = "off" // off (0) – GST_RPI_CAM_SRC_DRC_LEVEL_OFF
	GstRpiCamSrcDrclevelLow GstRpiCamSrcDrclevel = "low" // low (1) – GST_RPI_CAM_SRC_DRC_LEVEL_LOW
	GstRpiCamSrcDrclevelMedium GstRpiCamSrcDrclevel = "medium" // medium (2) – GST_RPI_CAM_SRC_DRC_LEVEL_MEDIUM
	GstRpiCamSrcDrclevelHigh GstRpiCamSrcDrclevel = "high" // high (3) – GST_RPI_CAM_SRC_DRC_LEVEL_HIGH
)

type GstRpiCamSrcExposureMeteringMode string

const (
	GstRpiCamSrcExposureMeteringModeAverage GstRpiCamSrcExposureMeteringMode = "average" // average (0) – GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_AVERAGE
	GstRpiCamSrcExposureMeteringModeSpot GstRpiCamSrcExposureMeteringMode = "spot" // spot (1) – GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_SPOT
	GstRpiCamSrcExposureMeteringModeBacklist GstRpiCamSrcExposureMeteringMode = "backlist" // backlist (2) – GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_BACKLIST
	GstRpiCamSrcExposureMeteringModeMatrix GstRpiCamSrcExposureMeteringMode = "matrix" // matrix (3) – GST_RPI_CAM_SRC_EXPOSURE_METERING_MODE_MATRIX
)

type GstRpiCamSrcExposureMode string

const (
	GstRpiCamSrcExposureModeOff GstRpiCamSrcExposureMode = "off" // off (0) – GST_RPI_CAM_SRC_EXPOSURE_MODE_OFF
	GstRpiCamSrcExposureModeAuto GstRpiCamSrcExposureMode = "auto" // auto (1) – GST_RPI_CAM_SRC_EXPOSURE_MODE_AUTO
	GstRpiCamSrcExposureModeNight GstRpiCamSrcExposureMode = "night" // night (2) – GST_RPI_CAM_SRC_EXPOSURE_MODE_NIGHT
	GstRpiCamSrcExposureModeNightpreview GstRpiCamSrcExposureMode = "nightpreview" // nightpreview (3) – GST_RPI_CAM_SRC_EXPOSURE_MODE_NIGHTPREVIEW
	GstRpiCamSrcExposureModeBacklight GstRpiCamSrcExposureMode = "backlight" // backlight (4) – GST_RPI_CAM_SRC_EXPOSURE_MODE_BACKLIGHT
	GstRpiCamSrcExposureModeSpotlight GstRpiCamSrcExposureMode = "spotlight" // spotlight (5) – GST_RPI_CAM_SRC_EXPOSURE_MODE_SPOTLIGHT
	GstRpiCamSrcExposureModeSports GstRpiCamSrcExposureMode = "sports" // sports (6) – GST_RPI_CAM_SRC_EXPOSURE_MODE_SPORTS
	GstRpiCamSrcExposureModeSnow GstRpiCamSrcExposureMode = "snow" // snow (7) – GST_RPI_CAM_SRC_EXPOSURE_MODE_SNOW
	GstRpiCamSrcExposureModeBeach GstRpiCamSrcExposureMode = "beach" // beach (8) – GST_RPI_CAM_SRC_EXPOSURE_MODE_BEACH
	GstRpiCamSrcExposureModeVerylong GstRpiCamSrcExposureMode = "verylong" // verylong (9) – GST_RPI_CAM_SRC_EXPOSURE_MODE_VERYLONG
	GstRpiCamSrcExposureModeFixedfps GstRpiCamSrcExposureMode = "fixedfps" // fixedfps (10) – GST_RPI_CAM_SRC_EXPOSURE_MODE_FIXEDFPS
	GstRpiCamSrcExposureModeAntishake GstRpiCamSrcExposureMode = "antishake" // antishake (11) – GST_RPI_CAM_SRC_EXPOSURE_MODE_ANTISHAKE
	GstRpiCamSrcExposureModeFireworks GstRpiCamSrcExposureMode = "fireworks" // fireworks (12) – GST_RPI_CAM_SRC_EXPOSURE_MODE_FIREWORKS
)

type GstRpiCamSrcImageEffect string

const (
	GstRpiCamSrcImageEffectNone GstRpiCamSrcImageEffect = "none" // none (0) – GST_RPI_CAM_SRC_IMAGEFX_NONE
	GstRpiCamSrcImageEffectNegative GstRpiCamSrcImageEffect = "negative" // negative (1) – GST_RPI_CAM_SRC_IMAGEFX_NEGATIVE
	GstRpiCamSrcImageEffectSolarize GstRpiCamSrcImageEffect = "solarize" // solarize (2) – GST_RPI_CAM_SRC_IMAGEFX_SOLARIZE
	GstRpiCamSrcImageEffectPosterize GstRpiCamSrcImageEffect = "posterize" // posterize (3) – GST_RPI_CAM_SRC_IMAGEFX_POSTERIZE
	GstRpiCamSrcImageEffectWhiteboard GstRpiCamSrcImageEffect = "whiteboard" // whiteboard (4) – GST_RPI_CAM_SRC_IMAGEFX_WHITEBOARD
	GstRpiCamSrcImageEffectBlackboard GstRpiCamSrcImageEffect = "blackboard" // blackboard (5) – GST_RPI_CAM_SRC_IMAGEFX_BLACKBOARD
	GstRpiCamSrcImageEffectSketch GstRpiCamSrcImageEffect = "sketch" // sketch (6) – GST_RPI_CAM_SRC_IMAGEFX_SKETCH
	GstRpiCamSrcImageEffectDenoise GstRpiCamSrcImageEffect = "denoise" // denoise (7) – GST_RPI_CAM_SRC_IMAGEFX_DENOISE
	GstRpiCamSrcImageEffectEmboss GstRpiCamSrcImageEffect = "emboss" // emboss (8) – GST_RPI_CAM_SRC_IMAGEFX_EMBOSS
	GstRpiCamSrcImageEffectOilpaint GstRpiCamSrcImageEffect = "oilpaint" // oilpaint (9) – GST_RPI_CAM_SRC_IMAGEFX_OILPAINT
	GstRpiCamSrcImageEffectHatch GstRpiCamSrcImageEffect = "hatch" // hatch (10) – GST_RPI_CAM_SRC_IMAGEFX_HATCH
	GstRpiCamSrcImageEffectGpen GstRpiCamSrcImageEffect = "gpen" // gpen (11) – GST_RPI_CAM_SRC_IMAGEFX_GPEN
	GstRpiCamSrcImageEffectPastel GstRpiCamSrcImageEffect = "pastel" // pastel (12) – GST_RPI_CAM_SRC_IMAGEFX_PASTEL
	GstRpiCamSrcImageEffectWatercolour GstRpiCamSrcImageEffect = "watercolour" // watercolour (13) – GST_RPI_CAM_SRC_IMAGEFX_WATERCOLOUR
	GstRpiCamSrcImageEffectFilm GstRpiCamSrcImageEffect = "film" // film (14) – GST_RPI_CAM_SRC_IMAGEFX_FILM
	GstRpiCamSrcImageEffectBlur GstRpiCamSrcImageEffect = "blur" // blur (15) – GST_RPI_CAM_SRC_IMAGEFX_BLUR
	GstRpiCamSrcImageEffectSaturation GstRpiCamSrcImageEffect = "saturation" // saturation (16) – GST_RPI_CAM_SRC_IMAGEFX_SATURATION
	GstRpiCamSrcImageEffectColourswap GstRpiCamSrcImageEffect = "colourswap" // colourswap (17) – GST_RPI_CAM_SRC_IMAGEFX_COLOURSWAP
	GstRpiCamSrcImageEffectWashedout GstRpiCamSrcImageEffect = "washedout" // washedout (18) – GST_RPI_CAM_SRC_IMAGEFX_WASHEDOUT
	GstRpiCamSrcImageEffectPosterise GstRpiCamSrcImageEffect = "posterise" // posterise (19) – GST_RPI_CAM_SRC_IMAGEFX_POSTERISE
	GstRpiCamSrcImageEffectColourpoint GstRpiCamSrcImageEffect = "colourpoint" // colourpoint (20) – GST_RPI_CAM_SRC_IMAGEFX_COLOURPOINT
	GstRpiCamSrcImageEffectColourbalance GstRpiCamSrcImageEffect = "colourbalance" // colourbalance (21) – GST_RPI_CAM_SRC_IMAGEFX_COLOURBALANCE
	GstRpiCamSrcImageEffectCartoon GstRpiCamSrcImageEffect = "cartoon" // cartoon (22) – GST_RPI_CAM_SRC_IMAGEFX_CARTOON
)

type GstRpiCamSrcIntraRefreshType string

const (
	GstRpiCamSrcIntraRefreshTypeNone GstRpiCamSrcIntraRefreshType = "none" // none (-1) – GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_NONE
	GstRpiCamSrcIntraRefreshTypeCyclic GstRpiCamSrcIntraRefreshType = "cyclic" // cyclic (0) – GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_CYCLIC
	GstRpiCamSrcIntraRefreshTypeAdaptive GstRpiCamSrcIntraRefreshType = "adaptive" // adaptive (1) – GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_ADAPTIVE
	GstRpiCamSrcIntraRefreshTypeBoth GstRpiCamSrcIntraRefreshType = "both" // both (2) – GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_BOTH
	GstRpiCamSrcIntraRefreshTypeCyclicRows GstRpiCamSrcIntraRefreshType = "cyclic-rows" // cyclic-rows (2130706433) – GST_RPI_CAM_SRC_INTRA_REFRESH_TYPE_CYCLIC_ROWS
)

type GstRpiCamSrcSensorMode string

const (
	GstRpiCamSrcSensorModeAutomatic GstRpiCamSrcSensorMode = "automatic" // automatic (0) – Automatic
	GstRpiCamSrcSensorMode1920X1080 GstRpiCamSrcSensorMode = "1920x1080" // 1920x1080 (1) – 1920x1080 16:9 1-30fps
	GstRpiCamSrcSensorMode2592X1944Fast GstRpiCamSrcSensorMode = "2592x1944-fast" // 2592x1944-fast (2) – 2592x1944 4:3 1-15fps / 3240x2464 15fps w/ v.2 board
	GstRpiCamSrcSensorMode2592X1944Slow GstRpiCamSrcSensorMode = "2592x1944-slow" // 2592x1944-slow (3) – 2592x1944 4:3 0.1666-1fps / 3240x2464 15fps w/ v.2 board
	GstRpiCamSrcSensorMode1296X972 GstRpiCamSrcSensorMode = "1296x972" // 1296x972 (4) – 1296x972 4:3 1-42fps
	GstRpiCamSrcSensorMode1296X730 GstRpiCamSrcSensorMode = "1296x730" // 1296x730 (5) – 1296x730 16:9 1-49fps
	GstRpiCamSrcSensorMode640X480Slow GstRpiCamSrcSensorMode = "640x480-slow" // 640x480-slow (6) – 640x480 4:3 42.1-60fps
	GstRpiCamSrcSensorMode640X480Fast GstRpiCamSrcSensorMode = "640x480-fast" // 640x480-fast (7) – 640x480 4:3 60.1-90fps
)

