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
)

type X265Enc struct {
	Element *gst.Element
}

func NewX265Enc() (*X265Enc, error) {
	e, err := gst.NewElement("x265enc")
	if err != nil {
		return nil, err
	}
	return &X265Enc{Element: e}, nil
}

func NewX265EncWithName(name string) (*X265Enc, error) {
	e, err := gst.NewElementWithName("x265enc", name)
	if err != nil {
		return nil, err
	}
	return &X265Enc{Element: e}, nil
}

// ----- Properties -----

// bitrate (uint)
//
// Bitrate in kbit/sec

func (e *X265Enc) GetBitrate() (uint, error) {
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

func (e *X265Enc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// key-int-max (int)
//
// Maximal distance between two key-frames (0 = x265 default / 250)

func (e *X265Enc) GetKeyIntMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("key-int-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *X265Enc) SetKeyIntMax(value int) error {
	return e.Element.SetProperty("key-int-max", value)
}

// log-level (GstX265LogLevel)
//
// x265 log level

func (e *X265Enc) GetLogLevel() (GstX265LogLevel, error) {
	var value GstX265LogLevel
	var ok bool
	v, err := e.Element.GetProperty("log-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX265LogLevel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX265LogLevel")
	}
	return value, nil
}

func (e *X265Enc) SetLogLevel(value GstX265LogLevel) error {
	e.Element.SetArg("log-level", string(value))
	return nil
}

// option-string (string)
//
// String of x265 options (overridden by element properties) in the format "key1=value1:key2=value2".

func (e *X265Enc) GetOptionString() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("option-string")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *X265Enc) SetOptionString(value string) error {
	return e.Element.SetProperty("option-string", value)
}

// qp (int)
//
// QP for P slices in (implied) CQP mode (-1 = disabled)

func (e *X265Enc) GetQp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *X265Enc) SetQp(value int) error {
	return e.Element.SetProperty("qp", value)
}

// speed-preset (GstX265SpeedPreset)
//
// Preset name for speed/quality tradeoff options

func (e *X265Enc) GetSpeedPreset() (GstX265SpeedPreset, error) {
	var value GstX265SpeedPreset
	var ok bool
	v, err := e.Element.GetProperty("speed-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX265SpeedPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX265SpeedPreset")
	}
	return value, nil
}

func (e *X265Enc) SetSpeedPreset(value GstX265SpeedPreset) error {
	e.Element.SetArg("speed-preset", string(value))
	return nil
}

// tune (GstX265Tune)
//
// Preset name for tuning options

func (e *X265Enc) GetTune() (GstX265Tune, error) {
	var value GstX265Tune
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstX265Tune)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstX265Tune")
	}
	return value, nil
}

func (e *X265Enc) SetTune(value GstX265Tune) error {
	e.Element.SetArg("tune", string(value))
	return nil
}

// ----- Constants -----

type GstX265SpeedPreset string

const (
	GstX265SpeedPresetNoPreset GstX265SpeedPreset = "No preset" // No preset (0) – No preset
	GstX265SpeedPresetUltrafast GstX265SpeedPreset = "ultrafast" // ultrafast (1) – ultrafast
	GstX265SpeedPresetSuperfast GstX265SpeedPreset = "superfast" // superfast (2) – superfast
	GstX265SpeedPresetVeryfast GstX265SpeedPreset = "veryfast" // veryfast (3) – veryfast
	GstX265SpeedPresetFaster GstX265SpeedPreset = "faster" // faster (4) – faster
	GstX265SpeedPresetFast GstX265SpeedPreset = "fast" // fast (5) – fast
	GstX265SpeedPresetMedium GstX265SpeedPreset = "medium" // medium (6) – medium
	GstX265SpeedPresetSlow GstX265SpeedPreset = "slow" // slow (7) – slow
	GstX265SpeedPresetSlower GstX265SpeedPreset = "slower" // slower (8) – slower
	GstX265SpeedPresetVeryslow GstX265SpeedPreset = "veryslow" // veryslow (9) – veryslow
	GstX265SpeedPresetPlacebo GstX265SpeedPreset = "placebo" // placebo (10) – placebo
)

type GstX265Tune string

const (
	GstX265TuneNoTunning GstX265Tune = "No tunning" // No tunning (0) – No tunning
	GstX265TunePsnr GstX265Tune = "psnr" // psnr (1) – psnr
	GstX265TuneSsim GstX265Tune = "ssim" // ssim (2) – ssim
	GstX265TuneGrain GstX265Tune = "grain" // grain (3) – grain
	GstX265TuneZerolatency GstX265Tune = "zerolatency" // zerolatency (4) – zerolatency
	GstX265TuneFastdecode GstX265Tune = "fastdecode" // fastdecode (5) – fastdecode
	GstX265TuneAnimation GstX265Tune = "animation" // animation (6) – animation
)

type GstX265LogLevel string

const (
	GstX265LogLevelNone GstX265LogLevel = "none" // none (-1) – No logging
	GstX265LogLevelError GstX265LogLevel = "error" // error (0) – Error
	GstX265LogLevelWarning GstX265LogLevel = "warning" // warning (1) – Warning
	GstX265LogLevelInfo GstX265LogLevel = "info" // info (2) – Info
	GstX265LogLevelDebug GstX265LogLevel = "debug" // debug (3) – Debug
	GstX265LogLevelFull GstX265LogLevel = "full" // full (4) – Full
)

