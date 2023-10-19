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

type Decklinkvideosrc struct {
	*base.GstPushSrc
}

func NewDecklinkvideosrc() (*Decklinkvideosrc, error) {
	e, err := gst.NewElement("decklinkvideosrc")
	if err != nil {
		return nil, err
	}
	return &Decklinkvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDecklinkvideosrcWithName(name string) (*Decklinkvideosrc, error) {
	e, err := gst.NewElementWithName("decklinkvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Decklinkvideosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// buffer-size (uint)
//
// Size of internal buffer in number of video frames

func (e *Decklinkvideosrc) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

// connection (GstDecklinkConnection)
//
// Video input connection to use

func (e *Decklinkvideosrc) GetConnection() (GstDecklinkConnection, error) {
	var value GstDecklinkConnection
	var ok bool
	v, err := e.Element.GetProperty("connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkConnection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkConnection")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetConnection(value GstDecklinkConnection) error {
	e.Element.SetArg("connection", string(value))
	return nil
}

// device-number (int)
//
// Output device instance to use

func (e *Decklinkvideosrc) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

// drop-no-signal-frames (bool)
//
// Drop frames that are marked as having no input signal

func (e *Decklinkvideosrc) GetDropNoSignalFrames() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-no-signal-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetDropNoSignalFrames(value bool) error {
	return e.Element.SetProperty("drop-no-signal-frames", value)
}

// hw-serial-number (string)
//
// The serial number (hardware ID) of the Decklink card

func (e *Decklinkvideosrc) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// mode (GstDecklinkModes)
//
// Video Mode to use for playback

func (e *Decklinkvideosrc) GetMode() (GstDecklinkModes, error) {
	var value GstDecklinkModes
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkModes")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetMode(value GstDecklinkModes) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// output-afd-bar (bool)
//
// Extract and output AFD/Bar as GstMeta (if present)

func (e *Decklinkvideosrc) GetOutputAfdBar() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-afd-bar")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetOutputAfdBar(value bool) error {
	return e.Element.SetProperty("output-afd-bar", value)
}

// output-cc (bool)
//
// Extract and output CC as GstMeta (if present)

func (e *Decklinkvideosrc) GetOutputCc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetOutputCc(value bool) error {
	return e.Element.SetProperty("output-cc", value)
}

// output-stream-time (bool)
//
// Output stream time directly instead of translating to pipeline clock

func (e *Decklinkvideosrc) GetOutputStreamTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-stream-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetOutputStreamTime(value bool) error {
	return e.Element.SetProperty("output-stream-time", value)
}

// persistent-id (int64)
//
// Decklink device to use. Higher priority than "device-number".
// BMDDeckLinkPersistentID is a device speciﬁc, 32-bit unique identiﬁer.
// It is stable even when the device is plugged in a diļ¬erent connector,
// across reboots, and when plugged into diļ¬erent computers.

func (e *Decklinkvideosrc) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

// profile (GstDecklinkProfileId)
//
// Specifies decklink profile to use.

func (e *Decklinkvideosrc) GetProfile() (GstDecklinkProfileId, error) {
	var value GstDecklinkProfileId
	var ok bool
	v, err := e.Element.GetProperty("profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkProfileId)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkProfileId")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetProfile(value GstDecklinkProfileId) error {
	e.Element.SetArg("profile", string(value))
	return nil
}

// signal (bool)
//
// True if there is a valid input signal available

func (e *Decklinkvideosrc) GetSignal() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// skip-first-time (uint64)
//
// Skip that much time of initial frames after starting

func (e *Decklinkvideosrc) GetSkipFirstTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("skip-first-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetSkipFirstTime(value uint64) error {
	return e.Element.SetProperty("skip-first-time", value)
}

// timecode-format (GstDecklinkTimecodeFormat)
//
// Timecode format type to use for input

func (e *Decklinkvideosrc) GetTimecodeFormat() (GstDecklinkTimecodeFormat, error) {
	var value GstDecklinkTimecodeFormat
	var ok bool
	v, err := e.Element.GetProperty("timecode-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkTimecodeFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkTimecodeFormat")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetTimecodeFormat(value GstDecklinkTimecodeFormat) error {
	e.Element.SetArg("timecode-format", string(value))
	return nil
}

// video-format (GstDecklinkVideoFormat)
//
// Video format type to use for input (Only use auto for mode=auto)

func (e *Decklinkvideosrc) GetVideoFormat() (GstDecklinkVideoFormat, error) {
	var value GstDecklinkVideoFormat
	var ok bool
	v, err := e.Element.GetProperty("video-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkVideoFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkVideoFormat")
	}
	return value, nil
}

func (e *Decklinkvideosrc) SetVideoFormat(value GstDecklinkVideoFormat) error {
	e.Element.SetArg("video-format", string(value))
	return nil
}

// ----- Constants -----

type GstDecklinkConnection string

const (
	GstDecklinkConnectionAuto GstDecklinkConnection = "auto" // auto (0) – Auto
	GstDecklinkConnectionSdi GstDecklinkConnection = "sdi" // sdi (1) – SDI
	GstDecklinkConnectionHdmi GstDecklinkConnection = "hdmi" // hdmi (2) – HDMI
	GstDecklinkConnectionOpticalSdi GstDecklinkConnection = "optical-sdi" // optical-sdi (3) – Optical SDI
	GstDecklinkConnectionComponent GstDecklinkConnection = "component" // component (4) – Component
	GstDecklinkConnectionComposite GstDecklinkConnection = "composite" // composite (5) – Composite
	GstDecklinkConnectionSvideo GstDecklinkConnection = "svideo" // svideo (6) – S-Video
)

type GstDecklinkModes string

const (
	GstDecklinkModesAuto GstDecklinkModes = "auto" // auto (0) – Automatic detection
	GstDecklinkModesNtsc GstDecklinkModes = "ntsc" // ntsc (1) – NTSC SD 60i
	GstDecklinkModesNtsc2398 GstDecklinkModes = "ntsc2398" // ntsc2398 (2) – NTSC SD 60i (24 fps)
	GstDecklinkModesPal GstDecklinkModes = "pal" // pal (3) – PAL SD 50i
	GstDecklinkModesNtscP GstDecklinkModes = "ntsc-p" // ntsc-p (4) – NTSC SD 60p
	GstDecklinkModesPalP GstDecklinkModes = "pal-p" // pal-p (5) – PAL SD 50p
	GstDecklinkModes1080P2398 GstDecklinkModes = "1080p2398" // 1080p2398 (6) – HD1080 23.98p
	GstDecklinkModes1080P24 GstDecklinkModes = "1080p24" // 1080p24 (7) – HD1080 24p
	GstDecklinkModes1080P25 GstDecklinkModes = "1080p25" // 1080p25 (8) – HD1080 25p
	GstDecklinkModes1080P2997 GstDecklinkModes = "1080p2997" // 1080p2997 (9) – HD1080 29.97p
	GstDecklinkModes1080P30 GstDecklinkModes = "1080p30" // 1080p30 (10) – HD1080 30p
	GstDecklinkModes1080I50 GstDecklinkModes = "1080i50" // 1080i50 (11) – HD1080 50i
	GstDecklinkModes1080I5994 GstDecklinkModes = "1080i5994" // 1080i5994 (12) – HD1080 59.94i
	GstDecklinkModes1080I60 GstDecklinkModes = "1080i60" // 1080i60 (13) – HD1080 60i
	GstDecklinkModes1080P50 GstDecklinkModes = "1080p50" // 1080p50 (14) – HD1080 50p
	GstDecklinkModes1080P5994 GstDecklinkModes = "1080p5994" // 1080p5994 (15) – HD1080 59.94p
	GstDecklinkModes1080P60 GstDecklinkModes = "1080p60" // 1080p60 (16) – HD1080 60p
	GstDecklinkModes720P50 GstDecklinkModes = "720p50" // 720p50 (17) – HD720 50p
	GstDecklinkModes720P5994 GstDecklinkModes = "720p5994" // 720p5994 (18) – HD720 59.94p
	GstDecklinkModes720P60 GstDecklinkModes = "720p60" // 720p60 (19) – HD720 60p
	GstDecklinkModes1556P2398 GstDecklinkModes = "1556p2398" // 1556p2398 (20) – 2k 23.98p
	GstDecklinkModes1556P24 GstDecklinkModes = "1556p24" // 1556p24 (21) – 2k 24p
	GstDecklinkModes1556P25 GstDecklinkModes = "1556p25" // 1556p25 (22) – 2k 25p
	GstDecklinkModes2Kdcip2398 GstDecklinkModes = "2kdcip2398" // 2kdcip2398 (23) – 2k dci 23.98p
	GstDecklinkModes2Kdcip24 GstDecklinkModes = "2kdcip24" // 2kdcip24 (24) – 2k dci 24p
	GstDecklinkModes2Kdcip25 GstDecklinkModes = "2kdcip25" // 2kdcip25 (25) – 2k dci 25p
	GstDecklinkModes2Kdcip2997 GstDecklinkModes = "2kdcip2997" // 2kdcip2997 (26) – 2k dci 29.97p
	GstDecklinkModes2Kdcip30 GstDecklinkModes = "2kdcip30" // 2kdcip30 (27) – 2k dci 30p
	GstDecklinkModes2Kdcip50 GstDecklinkModes = "2kdcip50" // 2kdcip50 (28) – 2k dci 50p
	GstDecklinkModes2Kdcip5994 GstDecklinkModes = "2kdcip5994" // 2kdcip5994 (29) – 2k dci 59.94p
	GstDecklinkModes2Kdcip60 GstDecklinkModes = "2kdcip60" // 2kdcip60 (30) – 2k dci 60p
	GstDecklinkModes2160P2398 GstDecklinkModes = "2160p2398" // 2160p2398 (31) – 4k 23.98p
	GstDecklinkModes2160P24 GstDecklinkModes = "2160p24" // 2160p24 (32) – 4k 24p
	GstDecklinkModes2160P25 GstDecklinkModes = "2160p25" // 2160p25 (33) – 4k 25p
	GstDecklinkModes2160P2997 GstDecklinkModes = "2160p2997" // 2160p2997 (34) – 4k 29.97p
	GstDecklinkModes2160P30 GstDecklinkModes = "2160p30" // 2160p30 (35) – 4k 30p
	GstDecklinkModes2160P50 GstDecklinkModes = "2160p50" // 2160p50 (36) – 4k 50p
	GstDecklinkModes2160P5994 GstDecklinkModes = "2160p5994" // 2160p5994 (37) – 4k 59.94p
	GstDecklinkModes2160P60 GstDecklinkModes = "2160p60" // 2160p60 (38) – 4k 60p
	GstDecklinkModesNtscWidescreen GstDecklinkModes = "ntsc-widescreen" // ntsc-widescreen (39) – NTSC SD 60i Widescreen
	GstDecklinkModesNtsc2398Widescreen GstDecklinkModes = "ntsc2398-widescreen" // ntsc2398-widescreen (40) – NTSC SD 60i Widescreen (24 fps)
	GstDecklinkModesPalWidescreen GstDecklinkModes = "pal-widescreen" // pal-widescreen (41) – PAL SD 50i Widescreen
	GstDecklinkModesNtscPWidescreen GstDecklinkModes = "ntsc-p-widescreen" // ntsc-p-widescreen (42) – NTSC SD 60p Widescreen
	GstDecklinkModesPalPWidescreen GstDecklinkModes = "pal-p-widescreen" // pal-p-widescreen (43) – PAL SD 50p Widescreen
	GstDecklinkModes4Kdcip2398 GstDecklinkModes = "4kdcip2398" // 4kdcip2398 (44) – 4k dci 23.98p
	GstDecklinkModes4Kdcip24 GstDecklinkModes = "4kdcip24" // 4kdcip24 (45) – 4k dci 24p
	GstDecklinkModes4Kdcip25 GstDecklinkModes = "4kdcip25" // 4kdcip25 (46) – 4k dci 25p
	GstDecklinkModes4Kdcip2997 GstDecklinkModes = "4kdcip2997" // 4kdcip2997 (47) – 4k dci 29.97p
	GstDecklinkModes4Kdcip30 GstDecklinkModes = "4kdcip30" // 4kdcip30 (48) – 4k dci 30p
	GstDecklinkModes4Kdcip50 GstDecklinkModes = "4kdcip50" // 4kdcip50 (49) – 4k dci 50p
	GstDecklinkModes4Kdcip5994 GstDecklinkModes = "4kdcip5994" // 4kdcip5994 (50) – 4k dci 59.94p
	GstDecklinkModes4Kdcip60 GstDecklinkModes = "4kdcip60" // 4kdcip60 (51) – 4k dci 60p
	GstDecklinkModes8Kp2398 GstDecklinkModes = "8kp2398" // 8kp2398 (52) – 8k 23.98p
	GstDecklinkModes8Kp24 GstDecklinkModes = "8kp24" // 8kp24 (53) – 8k 24p
	GstDecklinkModes8Kp25 GstDecklinkModes = "8kp25" // 8kp25 (54) – 8k 25p
	GstDecklinkModes8Kp2997 GstDecklinkModes = "8kp2997" // 8kp2997 (55) – 8k 29.97p
	GstDecklinkModes8Kp30 GstDecklinkModes = "8kp30" // 8kp30 (56) – 8k 30p
	GstDecklinkModes8Kp50 GstDecklinkModes = "8kp50" // 8kp50 (57) – 8k 50p
	GstDecklinkModes8Kp5994 GstDecklinkModes = "8kp5994" // 8kp5994 (58) – 8k 59.94p
	GstDecklinkModes8Kp60 GstDecklinkModes = "8kp60" // 8kp60 (59) – 8k 60p
	GstDecklinkModes8Kdcip2398 GstDecklinkModes = "8kdcip2398" // 8kdcip2398 (60) – 8k dci 23.98p
	GstDecklinkModes8Kdcip24 GstDecklinkModes = "8kdcip24" // 8kdcip24 (61) – 8k dci 24p
	GstDecklinkModes8Kdcip25 GstDecklinkModes = "8kdcip25" // 8kdcip25 (62) – 8k dci 25p
	GstDecklinkModes8Kdcip2997 GstDecklinkModes = "8kdcip2997" // 8kdcip2997 (63) – 8k dci 29.97p
	GstDecklinkModes8Kdcip30 GstDecklinkModes = "8kdcip30" // 8kdcip30 (64) – 8k dci 30p
	GstDecklinkModes8Kdcip50 GstDecklinkModes = "8kdcip50" // 8kdcip50 (65) – 8k dci 50p
	GstDecklinkModes8Kdcip5994 GstDecklinkModes = "8kdcip5994" // 8kdcip5994 (66) – 8k dci 59.94p
	GstDecklinkModes8Kdcip60 GstDecklinkModes = "8kdcip60" // 8kdcip60 (67) – 8k dci 60p
)

type GstDecklinkProfileId string

const (
	GstDecklinkProfileIdDefault GstDecklinkProfileId = "default" // default (0) – Default, don't change profile
	GstDecklinkProfileIdOneSubDeviceFull GstDecklinkProfileId = "one-sub-device-full" // one-sub-device-full (1) – One sub-device, Full-Duplex
	GstDecklinkProfileIdOneSubDeviceHalf GstDecklinkProfileId = "one-sub-device-half" // one-sub-device-half (2) – One sub-device, Half-Duplex
	GstDecklinkProfileIdTwoSubDevicesFull GstDecklinkProfileId = "two-sub-devices-full" // two-sub-devices-full (3) – Two sub-devices, Full-Duplex
	GstDecklinkProfileIdTwoSubDevicesHalf GstDecklinkProfileId = "two-sub-devices-half" // two-sub-devices-half (4) – Two sub-devices, Half-Duplex
	GstDecklinkProfileIdFourSubDevicesHalf GstDecklinkProfileId = "four-sub-devices-half" // four-sub-devices-half (5) – Four sub-devices, Half-Duplex
)

type GstDecklinkTimecodeFormat string

const (
	GstDecklinkTimecodeFormatRp188Vitc1 GstDecklinkTimecodeFormat = "rp188vitc1" // rp188vitc1 (0) – bmdTimecodeRP188VITC1
	GstDecklinkTimecodeFormatRp188Vitc2 GstDecklinkTimecodeFormat = "rp188vitc2" // rp188vitc2 (1) – bmdTimecodeRP188VITC2
	GstDecklinkTimecodeFormatRp188Ltc GstDecklinkTimecodeFormat = "rp188ltc" // rp188ltc (2) – bmdTimecodeRP188LTC
	GstDecklinkTimecodeFormatRp188Any GstDecklinkTimecodeFormat = "rp188any" // rp188any (3) – bmdTimecodeRP188Any
	GstDecklinkTimecodeFormatVitc GstDecklinkTimecodeFormat = "vitc" // vitc (4) – bmdTimecodeVITC
	GstDecklinkTimecodeFormatVitcfield2 GstDecklinkTimecodeFormat = "vitcfield2" // vitcfield2 (5) – bmdTimecodeVITCField2
	GstDecklinkTimecodeFormatSerial GstDecklinkTimecodeFormat = "serial" // serial (6) – bmdTimecodeSerial
)

type GstDecklinkVideoFormat string

const (
	GstDecklinkVideoFormatAuto GstDecklinkVideoFormat = "auto" // auto (0) – Auto
	GstDecklinkVideoFormat8BitYuv GstDecklinkVideoFormat = "8bit-yuv" // 8bit-yuv (1) – bmdFormat8BitYUV
	GstDecklinkVideoFormat10BitYuv GstDecklinkVideoFormat = "10bit-yuv" // 10bit-yuv (2) – bmdFormat10BitYUV
	GstDecklinkVideoFormat8BitArgb GstDecklinkVideoFormat = "8bit-argb" // 8bit-argb (3) – bmdFormat8BitARGB
	GstDecklinkVideoFormat8BitBgra GstDecklinkVideoFormat = "8bit-bgra" // 8bit-bgra (4) – bmdFormat8BitBGRA
	GstDecklinkVideoFormat10BitRgb GstDecklinkVideoFormat = "10bit-rgb" // 10bit-rgb (5) – bmdFormat10BitRGB
)

