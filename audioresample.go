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

type Audioresample struct {
	*base.GstBaseTransform
}

func NewAudioresample() (*Audioresample, error) {
	e, err := gst.NewElement("audioresample")
	if err != nil {
		return nil, err
	}
	return &Audioresample{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioresampleWithName(name string) (*Audioresample, error) {
	e, err := gst.NewElementWithName("audioresample", name)
	if err != nil {
		return nil, err
	}
	return &Audioresample{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// quality (int)
//
// Resample quality with 0 being the lowest and 10 being the best

func (e *Audioresample) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Audioresample) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

// resample-method (GstAudioResamplerMethod)
//
// What resample method to use

func (e *Audioresample) GetResampleMethod() (GstAudioResamplerMethod, error) {
	var value GstAudioResamplerMethod
	var ok bool
	v, err := e.Element.GetProperty("resample-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioResamplerMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioResamplerMethod")
	}
	return value, nil
}

func (e *Audioresample) SetResampleMethod(value GstAudioResamplerMethod) error {
	e.Element.SetArg("resample-method", string(value))
	return nil
}

// sinc-filter-auto-threshold (uint)
//
// Memory usage threshold to use if sinc filter mode is AUTO, given in bytes

func (e *Audioresample) GetSincFilterAutoThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sinc-filter-auto-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Audioresample) SetSincFilterAutoThreshold(value uint) error {
	return e.Element.SetProperty("sinc-filter-auto-threshold", value)
}

// sinc-filter-interpolation (GstAudioResamplerFilterInterpolation)
//
// How to interpolate the sinc filter table

func (e *Audioresample) GetSincFilterInterpolation() (GstAudioResamplerFilterInterpolation, error) {
	var value GstAudioResamplerFilterInterpolation
	var ok bool
	v, err := e.Element.GetProperty("sinc-filter-interpolation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioResamplerFilterInterpolation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioResamplerFilterInterpolation")
	}
	return value, nil
}

func (e *Audioresample) SetSincFilterInterpolation(value GstAudioResamplerFilterInterpolation) error {
	e.Element.SetArg("sinc-filter-interpolation", string(value))
	return nil
}

// sinc-filter-mode (GstAudioResamplerFilterMode)
//
// What sinc filter table mode to use

func (e *Audioresample) GetSincFilterMode() (GstAudioResamplerFilterMode, error) {
	var value GstAudioResamplerFilterMode
	var ok bool
	v, err := e.Element.GetProperty("sinc-filter-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioResamplerFilterMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioResamplerFilterMode")
	}
	return value, nil
}

func (e *Audioresample) SetSincFilterMode(value GstAudioResamplerFilterMode) error {
	e.Element.SetArg("sinc-filter-mode", string(value))
	return nil
}

// ----- Constants -----

type GstAudioResamplerFilterInterpolation string

const (
	GstAudioResamplerFilterInterpolationNone GstAudioResamplerFilterInterpolation = "none" // none (0) – GST_AUDIO_RESAMPLER_FILTER_INTERPOLATION_NONE
	GstAudioResamplerFilterInterpolationLinear GstAudioResamplerFilterInterpolation = "linear" // linear (1) – GST_AUDIO_RESAMPLER_FILTER_INTERPOLATION_LINEAR
	GstAudioResamplerFilterInterpolationCubic GstAudioResamplerFilterInterpolation = "cubic" // cubic (2) – GST_AUDIO_RESAMPLER_FILTER_INTERPOLATION_CUBIC
)

type GstAudioResamplerFilterMode string

const (
	GstAudioResamplerFilterModeInterpolated GstAudioResamplerFilterMode = "interpolated" // interpolated (0) – GST_AUDIO_RESAMPLER_FILTER_MODE_INTERPOLATED
	GstAudioResamplerFilterModeFull GstAudioResamplerFilterMode = "full" // full (1) – GST_AUDIO_RESAMPLER_FILTER_MODE_FULL
	GstAudioResamplerFilterModeAuto GstAudioResamplerFilterMode = "auto" // auto (2) – GST_AUDIO_RESAMPLER_FILTER_MODE_AUTO
)

type GstAudioResamplerMethod string

const (
	GstAudioResamplerMethodNearest GstAudioResamplerMethod = "nearest" // nearest (0) – GST_AUDIO_RESAMPLER_METHOD_NEAREST
	GstAudioResamplerMethodLinear GstAudioResamplerMethod = "linear" // linear (1) – GST_AUDIO_RESAMPLER_METHOD_LINEAR
	GstAudioResamplerMethodCubic GstAudioResamplerMethod = "cubic" // cubic (2) – GST_AUDIO_RESAMPLER_METHOD_CUBIC
	GstAudioResamplerMethodBlackmanNuttall GstAudioResamplerMethod = "blackman-nuttall" // blackman-nuttall (3) – GST_AUDIO_RESAMPLER_METHOD_BLACKMAN_NUTTALL
	GstAudioResamplerMethodKaiser GstAudioResamplerMethod = "kaiser" // kaiser (4) – GST_AUDIO_RESAMPLER_METHOD_KAISER
)

