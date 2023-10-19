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

type Audioconvert struct {
	*base.GstBaseTransform
}

func NewAudioconvert() (*Audioconvert, error) {
	e, err := gst.NewElement("audioconvert")
	if err != nil {
		return nil, err
	}
	return &Audioconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioconvertWithName(name string) (*Audioconvert, error) {
	e, err := gst.NewElementWithName("audioconvert", name)
	if err != nil {
		return nil, err
	}
	return &Audioconvert{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// dithering (GstAudioDitherMethod)
//
// Selects between different dithering methods.

func (e *Audioconvert) GetDithering() (interface{}, error) {
	return e.Element.GetProperty("dithering")
}

func (e *Audioconvert) SetDithering(value interface{}) error {
	return e.Element.SetProperty("dithering", value)
}

// dithering-threshold (uint)
//
// Threshold for the output bit depth at/below which to apply dithering.

func (e *Audioconvert) GetDitheringThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dithering-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Audioconvert) SetDitheringThreshold(value uint) error {
	return e.Element.SetProperty("dithering-threshold", value)
}

// mix-matrix (GstValueArray)
//
// Transformation matrix for input/output channels

func (e *Audioconvert) GetMixMatrix() (interface{}, error) {
	return e.Element.GetProperty("mix-matrix")
}

func (e *Audioconvert) SetMixMatrix(value interface{}) error {
	return e.Element.SetProperty("mix-matrix", value)
}

// noise-shaping (GstAudioNoiseShapingMethod)
//
// Selects between different noise shaping methods.

func (e *Audioconvert) GetNoiseShaping() (interface{}, error) {
	return e.Element.GetProperty("noise-shaping")
}

func (e *Audioconvert) SetNoiseShaping(value interface{}) error {
	return e.Element.SetProperty("noise-shaping", value)
}

