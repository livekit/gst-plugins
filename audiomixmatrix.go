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

type Audiomixmatrix struct {
	*base.GstBaseTransform
}

func NewAudiomixmatrix() (*Audiomixmatrix, error) {
	e, err := gst.NewElement("audiomixmatrix")
	if err != nil {
		return nil, err
	}
	return &Audiomixmatrix{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiomixmatrixWithName(name string) (*Audiomixmatrix, error) {
	e, err := gst.NewElementWithName("audiomixmatrix", name)
	if err != nil {
		return nil, err
	}
	return &Audiomixmatrix{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// channel-mask (uint64)
//
// Output channel mask (-1 means "default for these channels")

func (e *Audiomixmatrix) GetChannelMask() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("channel-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audiomixmatrix) SetChannelMask(value uint64) error {
	return e.Element.SetProperty("channel-mask", value)
}

// in-channels (uint)
//
// How many audio channels we have on the input side

func (e *Audiomixmatrix) GetInChannels() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("in-channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Audiomixmatrix) SetInChannels(value uint) error {
	return e.Element.SetProperty("in-channels", value)
}

// matrix (GstValueArray)
//
// Transformation matrix for input/output channels

func (e *Audiomixmatrix) GetMatrix() (interface{}, error) {
	return e.Element.GetProperty("matrix")
}

func (e *Audiomixmatrix) SetMatrix(value interface{}) error {
	return e.Element.SetProperty("matrix", value)
}

// mode (GstAudioMixMatrixModeType)
//
// Whether to auto-negotiate input/output channels and matrix

func (e *Audiomixmatrix) GetMode() (GstAudioMixMatrixModeType, error) {
	var value GstAudioMixMatrixModeType
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAudioMixMatrixModeType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAudioMixMatrixModeType")
	}
	return value, nil
}

func (e *Audiomixmatrix) SetMode(value GstAudioMixMatrixModeType) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// out-channels (uint)
//
// How many audio channels we have on the output side

func (e *Audiomixmatrix) GetOutChannels() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("out-channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Audiomixmatrix) SetOutChannels(value uint) error {
	return e.Element.SetProperty("out-channels", value)
}

// ----- Constants -----

type GstAudioMixMatrixModeType string

const (
	GstAudioMixMatrixModeTypeManual GstAudioMixMatrixModeType = "manual" // manual (0) – Manual mode: please specify input/output channels and transformation matrix
	GstAudioMixMatrixModeTypeFirstChannels GstAudioMixMatrixModeType = "first-channels" // first-channels (1) – First channels mode: input/output channels are auto-negotiated, transformation matrix is a truncated identity matrix
)

