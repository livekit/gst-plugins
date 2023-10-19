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

type V4L2Sink struct {
	*base.GstBaseSink
}

func NewV4L2Sink() (*V4L2Sink, error) {
	e, err := gst.NewElement("v4l2sink")
	if err != nil {
		return nil, err
	}
	return &V4L2Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewV4L2SinkWithName(name string) (*V4L2Sink, error) {
	e, err := gst.NewElementWithName("v4l2sink", name)
	if err != nil {
		return nil, err
	}
	return &V4L2Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// brightness (int)
//
// Picture brightness, or more precisely, the black level

func (e *V4L2Sink) GetBrightness() (int, error) {
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

func (e *V4L2Sink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (int)
//
// Picture contrast or luma gain

func (e *V4L2Sink) GetContrast() (int, error) {
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

func (e *V4L2Sink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// crop-height (uint)
//
// The height of the video crop; default is equal to negotiated image height

func (e *V4L2Sink) GetCropHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *V4L2Sink) SetCropHeight(value uint) error {
	return e.Element.SetProperty("crop-height", value)
}

// crop-left (int)
//
// The leftmost (x) coordinate of the video crop; top left corner of image is 0,0

func (e *V4L2Sink) GetCropLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("crop-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *V4L2Sink) SetCropLeft(value int) error {
	return e.Element.SetProperty("crop-left", value)
}

// crop-top (int)
//
// The topmost (y) coordinate of the video crop; top left corner of image is 0,0

func (e *V4L2Sink) GetCropTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("crop-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *V4L2Sink) SetCropTop(value int) error {
	return e.Element.SetProperty("crop-top", value)
}

// crop-width (uint)
//
// The width of the video crop; default is equal to negotiated image width

func (e *V4L2Sink) GetCropWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("crop-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *V4L2Sink) SetCropWidth(value uint) error {
	return e.Element.SetProperty("crop-width", value)
}

// device (string)
//
// Device location

func (e *V4L2Sink) GetDevice() (string, error) {
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

func (e *V4L2Sink) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-fd (int)
//
// File descriptor of the device

func (e *V4L2Sink) GetDeviceFd() (int, error) {
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

// device-name (string)
//
// Name of the device

func (e *V4L2Sink) GetDeviceName() (string, error) {
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

// extra-controls (GstStructure)
//
// Extra v4l2 controls (CIDs) for the device

func (e *V4L2Sink) GetExtraControls() (interface{}, error) {
	return e.Element.GetProperty("extra-controls")
}

func (e *V4L2Sink) SetExtraControls(value interface{}) error {
	return e.Element.SetProperty("extra-controls", value)
}

// flags (GstV4l2DeviceTypeFlags)
//
// Device type flags

func (e *V4L2Sink) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

// force-aspect-ratio (bool)
//
// When enabled, the pixel aspect ratio will be enforced

func (e *V4L2Sink) GetForceAspectRatio() (bool, error) {
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

func (e *V4L2Sink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// hue (int)
//
// Hue or color balance

func (e *V4L2Sink) GetHue() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *V4L2Sink) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

// io-mode (GstV4l2IOMode)
//
// I/O mode

func (e *V4L2Sink) GetIoMode() (interface{}, error) {
	return e.Element.GetProperty("io-mode")
}

func (e *V4L2Sink) SetIoMode(value interface{}) error {
	return e.Element.SetProperty("io-mode", value)
}

// norm (GstV4L2TvNorms)
//
// video standard

func (e *V4L2Sink) GetNorm() (interface{}, error) {
	return e.Element.GetProperty("norm")
}

func (e *V4L2Sink) SetNorm(value interface{}) error {
	return e.Element.SetProperty("norm", value)
}

// overlay-height (uint)
//
// The height of the video overlay; default is equal to negotiated image height

func (e *V4L2Sink) GetOverlayHeight() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("overlay-height")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *V4L2Sink) SetOverlayHeight(value uint) error {
	return e.Element.SetProperty("overlay-height", value)
}

// overlay-left (int)
//
// The leftmost (x) coordinate of the video overlay; top left corner of screen is 0,0

func (e *V4L2Sink) GetOverlayLeft() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *V4L2Sink) SetOverlayLeft(value int) error {
	return e.Element.SetProperty("overlay-left", value)
}

// overlay-top (int)
//
// The topmost (y) coordinate of the video overlay; top left corner of screen is 0,0

func (e *V4L2Sink) GetOverlayTop() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("overlay-top")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *V4L2Sink) SetOverlayTop(value int) error {
	return e.Element.SetProperty("overlay-top", value)
}

// overlay-width (uint)
//
// The width of the video overlay; default is equal to negotiated image width

func (e *V4L2Sink) GetOverlayWidth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("overlay-width")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *V4L2Sink) SetOverlayWidth(value uint) error {
	return e.Element.SetProperty("overlay-width", value)
}

// pixel-aspect-ratio (string)
//
// Overwrite the pixel aspect ratio of the device

func (e *V4L2Sink) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *V4L2Sink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// saturation (int)
//
// Picture color saturation or chroma gain

func (e *V4L2Sink) GetSaturation() (int, error) {
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

func (e *V4L2Sink) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

