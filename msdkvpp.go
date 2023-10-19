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

type Msdkvpp struct {
	*base.GstBaseTransform
}

func NewMsdkvpp() (*Msdkvpp, error) {
	e, err := gst.NewElement("msdkvpp")
	if err != nil {
		return nil, err
	}
	return &Msdkvpp{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewMsdkvppWithName(name string) (*Msdkvpp, error) {
	e, err := gst.NewElementWithName("msdkvpp", name)
	if err != nil {
		return nil, err
	}
	return &Msdkvpp{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// async-depth (uint)
//
// Depth of asynchronous pipeline

func (e *Msdkvpp) GetAsyncDepth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("async-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetAsyncDepth(value uint) error {
	return e.Element.SetProperty("async-depth", value)
}

// brightness (float32)
//
// The Brightness of the video

func (e *Msdkvpp) GetBrightness() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Msdkvpp) SetBrightness(value float32) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (float32)
//
// The Contrast of the video

func (e *Msdkvpp) GetContrast() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Msdkvpp) SetContrast(value float32) error {
	return e.Element.SetProperty("contrast", value)
}

// crop-bottom (uint)
//
// Pixels to crop at bottom

func (e *Msdkvpp) GetCropBottom() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-bottom")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetCropBottom(value uint) error {
	return e.Element.SetProperty("crop-bottom", value)
}

// crop-left (uint)
//
// Pixels to crop at left

func (e *Msdkvpp) GetCropLeft() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetCropLeft(value uint) error {
	return e.Element.SetProperty("crop-left", value)
}

// crop-right (uint)
//
// Pixels to crop at right

func (e *Msdkvpp) GetCropRight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetCropRight(value uint) error {
	return e.Element.SetProperty("crop-right", value)
}

// crop-top (uint)
//
// Pixels to crop at top

func (e *Msdkvpp) GetCropTop() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetCropTop(value uint) error {
	return e.Element.SetProperty("crop-top", value)
}

// deinterlace-method (GstMsdkVppdeinterlaceMethod)
//
// Deinterlace method to use

func (e *Msdkvpp) GetDeinterlaceMethod() (GstMsdkVppdeinterlaceMethod, error) {
	var value GstMsdkVppdeinterlaceMethod
	var ok bool
	v, err := e.Element.GetProperty("deinterlace-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVppdeinterlaceMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVppdeinterlaceMethod")
	}
	return value, nil
}

func (e *Msdkvpp) SetDeinterlaceMethod(value GstMsdkVppdeinterlaceMethod) error {
	e.Element.SetArg("deinterlace-method", string(value))
	return nil
}

// deinterlace-mode (GstMsdkVppdeinterlaceMode)
//
// Deinterlace mode to use

func (e *Msdkvpp) GetDeinterlaceMode() (GstMsdkVppdeinterlaceMode, error) {
	var value GstMsdkVppdeinterlaceMode
	var ok bool
	v, err := e.Element.GetProperty("deinterlace-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVppdeinterlaceMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVppdeinterlaceMode")
	}
	return value, nil
}

func (e *Msdkvpp) SetDeinterlaceMode(value GstMsdkVppdeinterlaceMode) error {
	e.Element.SetArg("deinterlace-mode", string(value))
	return nil
}

// denoise (uint)
//
// Denoising Factor

func (e *Msdkvpp) GetDenoise() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("denoise")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetDenoise(value uint) error {
	return e.Element.SetProperty("denoise", value)
}

// detail (uint)
//
// The factor of detail/edge enhancement filter algorithm

func (e *Msdkvpp) GetDetail() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("detail")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Msdkvpp) SetDetail(value uint) error {
	return e.Element.SetProperty("detail", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Msdkvpp) GetForceAspectRatio() (bool, error) {
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

func (e *Msdkvpp) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// frc-algorithm (GstMsdkVppfrcAlgorithm)
//
// The Framerate Control Alogorithm to use

func (e *Msdkvpp) GetFrcAlgorithm() (GstMsdkVppfrcAlgorithm, error) {
	var value GstMsdkVppfrcAlgorithm
	var ok bool
	v, err := e.Element.GetProperty("frc-algorithm")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVppfrcAlgorithm)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVppfrcAlgorithm")
	}
	return value, nil
}

func (e *Msdkvpp) SetFrcAlgorithm(value GstMsdkVppfrcAlgorithm) error {
	e.Element.SetArg("frc-algorithm", string(value))
	return nil
}

// hardware (bool)
//
// Enable hardware VPP

func (e *Msdkvpp) GetHardware() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hardware")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkvpp) SetHardware(value bool) error {
	return e.Element.SetProperty("hardware", value)
}

// hdr-tone-mapping (bool)

func (e *Msdkvpp) GetHdrToneMapping() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hdr-tone-mapping")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Msdkvpp) SetHdrToneMapping(value bool) error {
	return e.Element.SetProperty("hdr-tone-mapping", value)
}

// hue (float32)
//
// The hue of the video

func (e *Msdkvpp) GetHue() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Msdkvpp) SetHue(value float32) error {
	return e.Element.SetProperty("hue", value)
}

// mirroring (GstMsdkVppmirroring)
//
// The Mirroring type (DEPRECATED, use video-direction instead)

func (e *Msdkvpp) GetMirroring() (GstMsdkVppmirroring, error) {
	var value GstMsdkVppmirroring
	var ok bool
	v, err := e.Element.GetProperty("mirroring")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVppmirroring)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVppmirroring")
	}
	return value, nil
}

func (e *Msdkvpp) SetMirroring(value GstMsdkVppmirroring) error {
	e.Element.SetArg("mirroring", string(value))
	return nil
}

// rotation (GstMsdkVpprotation)
//
// Rotation Angle (DEPRECATED, use video-direction instead)

func (e *Msdkvpp) GetRotation() (GstMsdkVpprotation, error) {
	var value GstMsdkVpprotation
	var ok bool
	v, err := e.Element.GetProperty("rotation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVpprotation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVpprotation")
	}
	return value, nil
}

func (e *Msdkvpp) SetRotation(value GstMsdkVpprotation) error {
	e.Element.SetArg("rotation", string(value))
	return nil
}

// saturation (float32)
//
// The Saturation of the video

func (e *Msdkvpp) GetSaturation() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Msdkvpp) SetSaturation(value float32) error {
	return e.Element.SetProperty("saturation", value)
}

// scaling-mode (GstMsdkVppscalingMode)
//
// The Scaling mode to use

func (e *Msdkvpp) GetScalingMode() (GstMsdkVppscalingMode, error) {
	var value GstMsdkVppscalingMode
	var ok bool
	v, err := e.Element.GetProperty("scaling-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMsdkVppscalingMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMsdkVppscalingMode")
	}
	return value, nil
}

func (e *Msdkvpp) SetScalingMode(value GstMsdkVppscalingMode) error {
	e.Element.SetArg("scaling-mode", string(value))
	return nil
}

// video-direction (GstVideoOrientationMethod)
//
// Video direction: rotation and flipping, it will override both mirroring & rotation properties if set explicitly

func (e *Msdkvpp) GetVideoDirection() (interface{}, error) {
	return e.Element.GetProperty("video-direction")
}

func (e *Msdkvpp) SetVideoDirection(value interface{}) error {
	return e.Element.SetProperty("video-direction", value)
}

// ----- Constants -----

type GstMsdkVppfrcAlgorithm string

const (
	GstMsdkVppfrcAlgorithmNone GstMsdkVppfrcAlgorithm = "none" // none (0) – No FrameRate Control algorithm
	GstMsdkVppfrcAlgorithmPreserveTs GstMsdkVppfrcAlgorithm = "preserve-ts" // preserve-ts (1) – Frame dropping/repetition, Preserve timestamp
	GstMsdkVppfrcAlgorithmDistributeTs GstMsdkVppfrcAlgorithm = "distribute-ts" // distribute-ts (2) – Frame dropping/repetition, Distribute timestamp
	GstMsdkVppfrcAlgorithmInterpolate GstMsdkVppfrcAlgorithm = "interpolate" // interpolate (4) – Frame interpolation
	GstMsdkVppfrcAlgorithmInterpolatePreserveTs GstMsdkVppfrcAlgorithm = "interpolate-preserve-ts" // interpolate-preserve-ts (5) – Frame interpolation, Preserve timestamp
	GstMsdkVppfrcAlgorithmInterpolateDistributeTs GstMsdkVppfrcAlgorithm = "interpolate-distribute-ts" // interpolate-distribute-ts (6) – Frame interpolation, Distribute timestamp
)

type GstMsdkVppmirroring string

const (
	GstMsdkVppmirroringDisable GstMsdkVppmirroring = "disable" // disable (0) – Disable mirroring
	GstMsdkVppmirroringHorizontal GstMsdkVppmirroring = "horizontal" // horizontal (1) – Horizontal Mirroring
	GstMsdkVppmirroringVertical GstMsdkVppmirroring = "vertical" // vertical (2) – Vertical Mirroring
)

type GstMsdkVpprotation string

const (
	GstMsdkVpprotation0 GstMsdkVpprotation = "0" // 0 (0) – Unrotated mode
	GstMsdkVpprotation90 GstMsdkVpprotation = "90" // 90 (90) – Rotated by 90°
	GstMsdkVpprotation180 GstMsdkVpprotation = "180" // 180 (180) – Rotated by 180°
	GstMsdkVpprotation270 GstMsdkVpprotation = "270" // 270 (270) – Rotated by 270°
)

type GstMsdkVppscalingMode string

const (
	GstMsdkVppscalingModeDisable GstMsdkVppscalingMode = "disable" // disable (0) – Default Scaling
	GstMsdkVppscalingModeLowpower GstMsdkVppscalingMode = "lowpower" // lowpower (1) – Lowpower Scaling
	GstMsdkVppscalingModeQuality GstMsdkVppscalingMode = "quality" // quality (2) – High Quality Scaling
)

type GstMsdkVppdeinterlaceMethod string

const (
	GstMsdkVppdeinterlaceMethodNone GstMsdkVppdeinterlaceMethod = "none" // none (0) – Disable deinterlacing
	GstMsdkVppdeinterlaceMethodBob GstMsdkVppdeinterlaceMethod = "bob" // bob (1) – Bob deinterlacing
	GstMsdkVppdeinterlaceMethodAdvanced GstMsdkVppdeinterlaceMethod = "advanced" // advanced (2) – Advanced deinterlacing (Motion adaptive)
	GstMsdkVppdeinterlaceMethodAdvancedNoRef GstMsdkVppdeinterlaceMethod = "advanced-no-ref " // advanced-no-ref  (11) – Advanced deinterlacing mode without using of reference frames
	GstMsdkVppdeinterlaceMethodAdvancedScd GstMsdkVppdeinterlaceMethod = "advanced-scd " // advanced-scd  (12) – Advanced deinterlacing mode with scene change detection
	GstMsdkVppdeinterlaceMethodFieldWeave GstMsdkVppdeinterlaceMethod = "field-weave" // field-weave (13) – Field weaving
)

type GstMsdkVppdeinterlaceMode string

const (
	GstMsdkVppdeinterlaceModeAuto GstMsdkVppdeinterlaceMode = "auto" // auto (0) – Auto detection
	GstMsdkVppdeinterlaceModeInterlaced GstMsdkVppdeinterlaceMode = "interlaced" // interlaced (1) – Force deinterlacing
	GstMsdkVppdeinterlaceModeDisabled GstMsdkVppdeinterlaceMode = "disabled" // disabled (2) – Never deinterlace
)

