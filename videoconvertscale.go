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

type Videoconvertscale struct {
	*base.GstBaseTransform
}

func NewVideoconvertscale() (*Videoconvertscale, error) {
	e, err := gst.NewElement("videoconvertscale")
	if err != nil {
		return nil, err
	}
	return &Videoconvertscale{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideoconvertscaleWithName(name string) (*Videoconvertscale, error) {
	e, err := gst.NewElementWithName("videoconvertscale", name)
	if err != nil {
		return nil, err
	}
	return &Videoconvertscale{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// add-borders (bool)
//
// Add black borders if necessary to keep the display aspect ratio

func (e *Videoconvertscale) GetAddBorders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-borders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videoconvertscale) SetAddBorders(value bool) error {
	return e.Element.SetProperty("add-borders", value)
}

// alpha-mode (GstVideoAlphaMode)
//
// Alpha Mode to use

func (e *Videoconvertscale) GetAlphaMode() (interface{}, error) {
	return e.Element.GetProperty("alpha-mode")
}

func (e *Videoconvertscale) SetAlphaMode(value interface{}) error {
	return e.Element.SetProperty("alpha-mode", value)
}

// alpha-value (float64)
//
// Alpha Value to use

func (e *Videoconvertscale) GetAlphaValue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("alpha-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videoconvertscale) SetAlphaValue(value float64) error {
	return e.Element.SetProperty("alpha-value", value)
}

// chroma-mode (GstVideoChromaMode)
//
// Chroma Resampling Mode

func (e *Videoconvertscale) GetChromaMode() (interface{}, error) {
	return e.Element.GetProperty("chroma-mode")
}

func (e *Videoconvertscale) SetChromaMode(value interface{}) error {
	return e.Element.SetProperty("chroma-mode", value)
}

// chroma-resampler (GstVideoResamplerMethod)
//
// Chroma resampler method

func (e *Videoconvertscale) GetChromaResampler() (interface{}, error) {
	return e.Element.GetProperty("chroma-resampler")
}

func (e *Videoconvertscale) SetChromaResampler(value interface{}) error {
	return e.Element.SetProperty("chroma-resampler", value)
}

// converter-config (GstStructure)
//
// A GstStructure describing the configuration that should be used. This
// configuration, if set, takes precedence over the other similar conversion
// properties.

func (e *Videoconvertscale) GetConverterConfig() (interface{}, error) {
	return e.Element.GetProperty("converter-config")
}

func (e *Videoconvertscale) SetConverterConfig(value interface{}) error {
	return e.Element.SetProperty("converter-config", value)
}

// dither (GstVideoDitherMethod)
//
// Apply dithering while converting

func (e *Videoconvertscale) GetDither() (interface{}, error) {
	return e.Element.GetProperty("dither")
}

func (e *Videoconvertscale) SetDither(value interface{}) error {
	return e.Element.SetProperty("dither", value)
}

// dither-quantization (uint)
//
// Quantizer to use

func (e *Videoconvertscale) GetDitherQuantization() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dither-quantization")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Videoconvertscale) SetDitherQuantization(value uint) error {
	return e.Element.SetProperty("dither-quantization", value)
}

// envelope (float64)
//
// Size of filter envelope

func (e *Videoconvertscale) GetEnvelope() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("envelope")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videoconvertscale) SetEnvelope(value float64) error {
	return e.Element.SetProperty("envelope", value)
}

// gamma-mode (GstVideoGammaMode)
//
// Gamma Conversion Mode

func (e *Videoconvertscale) GetGammaMode() (interface{}, error) {
	return e.Element.GetProperty("gamma-mode")
}

func (e *Videoconvertscale) SetGammaMode(value interface{}) error {
	return e.Element.SetProperty("gamma-mode", value)
}

// matrix-mode (GstVideoMatrixMode)
//
// Matrix Conversion Mode

func (e *Videoconvertscale) GetMatrixMode() (interface{}, error) {
	return e.Element.GetProperty("matrix-mode")
}

func (e *Videoconvertscale) SetMatrixMode(value interface{}) error {
	return e.Element.SetProperty("matrix-mode", value)
}

// method (GstVideoScaleMethod)
//
// method

func (e *Videoconvertscale) GetMethod() (interface{}, error) {
	return e.Element.GetProperty("method")
}

func (e *Videoconvertscale) SetMethod(value interface{}) error {
	return e.Element.SetProperty("method", value)
}

// n-threads (uint)
//
// Maximum number of threads to use

func (e *Videoconvertscale) GetNThreads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("n-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Videoconvertscale) SetNThreads(value uint) error {
	return e.Element.SetProperty("n-threads", value)
}

// primaries-mode (GstVideoPrimariesMode)
//
// Primaries Conversion Mode

func (e *Videoconvertscale) GetPrimariesMode() (interface{}, error) {
	return e.Element.GetProperty("primaries-mode")
}

func (e *Videoconvertscale) SetPrimariesMode(value interface{}) error {
	return e.Element.SetProperty("primaries-mode", value)
}

// sharpen (float64)
//
// Sharpening

func (e *Videoconvertscale) GetSharpen() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sharpen")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videoconvertscale) SetSharpen(value float64) error {
	return e.Element.SetProperty("sharpen", value)
}

// sharpness (float64)
//
// Sharpness of filter

func (e *Videoconvertscale) GetSharpness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("sharpness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videoconvertscale) SetSharpness(value float64) error {
	return e.Element.SetProperty("sharpness", value)
}

