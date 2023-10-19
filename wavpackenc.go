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

type Wavpackenc struct {
	Element *gst.Element
}

func NewWavpackenc() (*Wavpackenc, error) {
	e, err := gst.NewElement("wavpackenc")
	if err != nil {
		return nil, err
	}
	return &Wavpackenc{Element: e}, nil
}

func NewWavpackencWithName(name string) (*Wavpackenc, error) {
	e, err := gst.NewElementWithName("wavpackenc", name)
	if err != nil {
		return nil, err
	}
	return &Wavpackenc{Element: e}, nil
}

// ----- Properties -----

// bitrate (uint)
//
// Try to encode with this average bitrate (bits/sec). This enables lossy encoding, values smaller than 24000 disable it again.

func (e *Wavpackenc) GetBitrate() (uint, error) {
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

func (e *Wavpackenc) SetBitrate(value uint) error {
	return e.Element.SetProperty("bitrate", value)
}

// bits-per-sample (float64)
//
// Try to encode with this amount of bits per sample. This enables lossy encoding, values smaller than 2.0 disable it again.

func (e *Wavpackenc) GetBitsPerSample() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bits-per-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Wavpackenc) SetBitsPerSample(value float64) error {
	return e.Element.SetProperty("bits-per-sample", value)
}

// correction-mode (GstWavpackEncCorrectionMode)
//
// Use this mode for the correction stream. Only works in lossy mode!

func (e *Wavpackenc) GetCorrectionMode() (GstWavpackEncCorrectionMode, error) {
	var value GstWavpackEncCorrectionMode
	var ok bool
	v, err := e.Element.GetProperty("correction-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWavpackEncCorrectionMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWavpackEncCorrectionMode")
	}
	return value, nil
}

func (e *Wavpackenc) SetCorrectionMode(value GstWavpackEncCorrectionMode) error {
	e.Element.SetArg("correction-mode", string(value))
	return nil
}

// extra-processing (uint)
//
// Use better but slower filters for better compression/quality.

func (e *Wavpackenc) GetExtraProcessing() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("extra-processing")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Wavpackenc) SetExtraProcessing(value uint) error {
	return e.Element.SetProperty("extra-processing", value)
}

// joint-stereo-mode (GstWavpackEncJsmode)
//
// Use this joint-stereo mode.

func (e *Wavpackenc) GetJointStereoMode() (GstWavpackEncJsmode, error) {
	var value GstWavpackEncJsmode
	var ok bool
	v, err := e.Element.GetProperty("joint-stereo-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWavpackEncJsmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWavpackEncJsmode")
	}
	return value, nil
}

func (e *Wavpackenc) SetJointStereoMode(value GstWavpackEncJsmode) error {
	e.Element.SetArg("joint-stereo-mode", string(value))
	return nil
}

// md5 (bool)
//
// Store MD5 hash of raw samples within the file.

func (e *Wavpackenc) GetMd5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Wavpackenc) SetMd5(value bool) error {
	return e.Element.SetProperty("md5", value)
}

// mode (GstWavpackEncMode)
//
// Speed versus compression tradeoff.

func (e *Wavpackenc) GetMode() (GstWavpackEncMode, error) {
	var value GstWavpackEncMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWavpackEncMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWavpackEncMode")
	}
	return value, nil
}

func (e *Wavpackenc) SetMode(value GstWavpackEncMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ----- Constants -----

type GstWavpackEncCorrectionMode string

const (
	GstWavpackEncCorrectionModeOff GstWavpackEncCorrectionMode = "off" // off (0) – Create no correction file
	GstWavpackEncCorrectionModeOn GstWavpackEncCorrectionMode = "on" // on (1) – Create correction file
	GstWavpackEncCorrectionModeOptimized GstWavpackEncCorrectionMode = "optimized" // optimized (2) – Create optimized correction file
)

type GstWavpackEncJsmode string

const (
	GstWavpackEncJsmodeAuto GstWavpackEncJsmode = "auto" // auto (0) – auto
	GstWavpackEncJsmodeLeftright GstWavpackEncJsmode = "leftright" // leftright (1) – left/right
	GstWavpackEncJsmodeMidside GstWavpackEncJsmode = "midside" // midside (2) – mid/side
)

type GstWavpackEncMode string

const (
	GstWavpackEncModeFast GstWavpackEncMode = "fast" // fast (1) – Fast Compression
	GstWavpackEncModeNormal GstWavpackEncMode = "normal" // normal (2) – Normal Compression
	GstWavpackEncModeHigh GstWavpackEncMode = "high" // high (3) – High Compression
	GstWavpackEncModeVeryhigh GstWavpackEncMode = "veryhigh" // veryhigh (4) – Very High Compression
)

