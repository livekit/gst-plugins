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

type V4L2Src struct {
	*base.GstPushSrc
}

func NewV4L2Src() (*V4L2Src, error) {
	e, err := gst.NewElement("v4l2src")
	if err != nil {
		return nil, err
	}
	return &V4L2Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewV4L2SrcWithName(name string) (*V4L2Src, error) {
	e, err := gst.NewElementWithName("v4l2src", name)
	if err != nil {
		return nil, err
	}
	return &V4L2Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// brightness (int)
//
// Picture brightness, or more precisely, the black level

func (e *V4L2Src) GetBrightness() (int, error) {
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

func (e *V4L2Src) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (int)
//
// Picture contrast or luma gain

func (e *V4L2Src) GetContrast() (int, error) {
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

func (e *V4L2Src) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// crop-bottom (uint)
//
// Number of pixels to crop from the bottom edge of captured video
// stream

func (e *V4L2Src) GetCropBottom() (uint, error) {
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

func (e *V4L2Src) SetCropBottom(value uint) error {
	return e.Element.SetProperty("crop-bottom", value)
}

// crop-bounds (GstValueArray)
//
// Crop bounding region.  All crop regions must lie within this region.
// The bounds are represented as a four element array, that descibes the
// [x, y, width, height] of the area.

// crop-left (uint)
//
// Number of pixels to crop from the left edge of captured video
// stream

func (e *V4L2Src) GetCropLeft() (uint, error) {
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

func (e *V4L2Src) SetCropLeft(value uint) error {
	return e.Element.SetProperty("crop-left", value)
}

// crop-right (uint)
//
// Number of pixels to crop from the right edge of captured video
// stream

func (e *V4L2Src) GetCropRight() (uint, error) {
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

func (e *V4L2Src) SetCropRight(value uint) error {
	return e.Element.SetProperty("crop-right", value)
}

// crop-top (uint)
//
// Number of pixels to crop from the top edge of captured video
// stream

func (e *V4L2Src) GetCropTop() (uint, error) {
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

func (e *V4L2Src) SetCropTop(value uint) error {
	return e.Element.SetProperty("crop-top", value)
}

// device (string)
//
// Device location

func (e *V4L2Src) GetDevice() (string, error) {
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

func (e *V4L2Src) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// device-fd (int)
//
// File descriptor of the device

func (e *V4L2Src) GetDeviceFd() (int, error) {
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

func (e *V4L2Src) GetDeviceName() (string, error) {
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
// Additional v4l2 controls for the device. The controls are identified
// by the control name (lowercase with '_' for any non-alphanumeric
// characters).

func (e *V4L2Src) GetExtraControls() (interface{}, error) {
	return e.Element.GetProperty("extra-controls")
}

func (e *V4L2Src) SetExtraControls(value interface{}) error {
	return e.Element.SetProperty("extra-controls", value)
}

// flags (GstV4l2DeviceTypeFlags)
//
// Device type flags

func (e *V4L2Src) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

// force-aspect-ratio (bool)
//
// When enabled, the pixel aspect ratio queried from the device or set
// with the pixel-aspect-ratio property will be enforced.

func (e *V4L2Src) GetForceAspectRatio() (bool, error) {
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

func (e *V4L2Src) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// hue (int)
//
// Hue or color balance

func (e *V4L2Src) GetHue() (int, error) {
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

func (e *V4L2Src) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

// io-mode (GstV4l2IOMode)
//
// IO Mode

func (e *V4L2Src) GetIoMode() (interface{}, error) {
	return e.Element.GetProperty("io-mode")
}

func (e *V4L2Src) SetIoMode(value interface{}) error {
	return e.Element.SetProperty("io-mode", value)
}

// norm (GstV4L2TvNorms)
//
// TV norm

func (e *V4L2Src) GetNorm() (interface{}, error) {
	return e.Element.GetProperty("norm")
}

func (e *V4L2Src) SetNorm(value interface{}) error {
	return e.Element.SetProperty("norm", value)
}

// pixel-aspect-ratio (string)
//
// The pixel aspect ratio of the device. This overwrites the pixel aspect
// ratio queried from the device.

func (e *V4L2Src) GetPixelAspectRatio() (string, error) {
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

func (e *V4L2Src) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// saturation (int)
//
// Picture color saturation or chroma gain

func (e *V4L2Src) GetSaturation() (int, error) {
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

func (e *V4L2Src) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

// ----- Constants -----

type GstV4L2DeviceTypeFlags string

const (
	GstV4L2DeviceTypeFlagsCapture GstV4L2DeviceTypeFlags = "capture" // capture (0x00000001) – Device supports video capture
	GstV4L2DeviceTypeFlagsOutput GstV4L2DeviceTypeFlags = "output" // output (0x00000002) – Device supports video playback
	GstV4L2DeviceTypeFlagsOverlay GstV4L2DeviceTypeFlags = "overlay" // overlay (0x00000004) – Device supports video overlay
	GstV4L2DeviceTypeFlagsVbiCapture GstV4L2DeviceTypeFlags = "vbi-capture" // vbi-capture (0x00000010) – Device supports the VBI capture
	GstV4L2DeviceTypeFlagsVbiOutput GstV4L2DeviceTypeFlags = "vbi-output" // vbi-output (0x00000020) – Device supports the VBI output
	GstV4L2DeviceTypeFlagsTuner GstV4L2DeviceTypeFlags = "tuner" // tuner (0x00010000) – Device has a tuner or modulator
	GstV4L2DeviceTypeFlagsAudio GstV4L2DeviceTypeFlags = "audio" // audio (0x00020000) – Device has audio inputs or outputs
)

type GstV4L2Iomode string

const (
	GstV4L2IomodeAuto GstV4L2Iomode = "auto" // auto (0) – GST_V4L2_IO_AUTO
	GstV4L2IomodeRw GstV4L2Iomode = "rw" // rw (1) – GST_V4L2_IO_RW
	GstV4L2IomodeMmap GstV4L2Iomode = "mmap" // mmap (2) – GST_V4L2_IO_MMAP
	GstV4L2IomodeUserptr GstV4L2Iomode = "userptr" // userptr (3) – GST_V4L2_IO_USERPTR
	GstV4L2IomodeDmabuf GstV4L2Iomode = "dmabuf" // dmabuf (4) – GST_V4L2_IO_DMABUF
	GstV4L2IomodeDmabufImport GstV4L2Iomode = "dmabuf-import" // dmabuf-import (5) – GST_V4L2_IO_DMABUF_IMPORT
)

type V4L2TvNorms string

const (
	V4L2TvNormsNone V4L2TvNorms = "none" // none (0) – none
	V4L2TvNormsNtsc V4L2TvNorms = "NTSC" // NTSC (45056) – NTSC
	V4L2TvNormsNtscM V4L2TvNorms = "NTSC-M" // NTSC-M (4096) – NTSC-M
	V4L2TvNormsNtscMJp V4L2TvNorms = "NTSC-M-JP" // NTSC-M-JP (8192) – NTSC-M-JP
	V4L2TvNormsNtscMKr V4L2TvNorms = "NTSC-M-KR" // NTSC-M-KR (32768) – NTSC-M-KR
	V4L2TvNormsNtsc443 V4L2TvNorms = "NTSC-443" // NTSC-443 (16384) – NTSC-443
	V4L2TvNormsPal V4L2TvNorms = "PAL" // PAL (255) – PAL
	V4L2TvNormsPalBg V4L2TvNorms = "PAL-BG" // PAL-BG (7) – PAL-BG
	V4L2TvNormsPalB V4L2TvNorms = "PAL-B" // PAL-B (1) – PAL-B
	V4L2TvNormsPalB1 V4L2TvNorms = "PAL-B1" // PAL-B1 (2) – PAL-B1
	V4L2TvNormsPalG V4L2TvNorms = "PAL-G" // PAL-G (4) – PAL-G
	V4L2TvNormsPalH V4L2TvNorms = "PAL-H" // PAL-H (8) – PAL-H
	V4L2TvNormsPalI V4L2TvNorms = "PAL-I" // PAL-I (16) – PAL-I
	V4L2TvNormsPalDk V4L2TvNorms = "PAL-DK" // PAL-DK (224) – PAL-DK
	V4L2TvNormsPalD V4L2TvNorms = "PAL-D" // PAL-D (32) – PAL-D
	V4L2TvNormsPalD1 V4L2TvNorms = "PAL-D1" // PAL-D1 (64) – PAL-D1
	V4L2TvNormsPalK V4L2TvNorms = "PAL-K" // PAL-K (128) – PAL-K
	V4L2TvNormsPalM V4L2TvNorms = "PAL-M" // PAL-M (256) – PAL-M
	V4L2TvNormsPalN V4L2TvNorms = "PAL-N" // PAL-N (512) – PAL-N
	V4L2TvNormsPalNc V4L2TvNorms = "PAL-Nc" // PAL-Nc (1024) – PAL-Nc
	V4L2TvNormsPal60 V4L2TvNorms = "PAL-60" // PAL-60 (2048) – PAL-60
	V4L2TvNormsSecam V4L2TvNorms = "SECAM" // SECAM (16711680) – SECAM
	V4L2TvNormsSecamB V4L2TvNorms = "SECAM-B" // SECAM-B (65536) – SECAM-B
	V4L2TvNormsSecamG V4L2TvNorms = "SECAM-G" // SECAM-G (262144) – SECAM-G
	V4L2TvNormsSecamH V4L2TvNorms = "SECAM-H" // SECAM-H (524288) – SECAM-H
	V4L2TvNormsSecamDk V4L2TvNorms = "SECAM-DK" // SECAM-DK (3276800) – SECAM-DK
	V4L2TvNormsSecamD V4L2TvNorms = "SECAM-D" // SECAM-D (131072) – SECAM-D
	V4L2TvNormsSecamK V4L2TvNorms = "SECAM-K" // SECAM-K (1048576) – SECAM-K
	V4L2TvNormsSecamK1 V4L2TvNorms = "SECAM-K1" // SECAM-K1 (2097152) – SECAM-K1
	V4L2TvNormsSecamL V4L2TvNorms = "SECAM-L" // SECAM-L (4194304) – SECAM-L
	V4L2TvNormsSecamLc V4L2TvNorms = "SECAM-Lc" // SECAM-Lc (8388608) – SECAM-Lc
)

