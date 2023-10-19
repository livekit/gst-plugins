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

type EqualizerNbands struct {
	*base.GstBaseTransform
}

func NewEqualizerNbands() (*EqualizerNbands, error) {
	e, err := gst.NewElement("equalizer-nbands")
	if err != nil {
		return nil, err
	}
	return &EqualizerNbands{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewEqualizerNbandsWithName(name string) (*EqualizerNbands, error) {
	e, err := gst.NewElementWithName("equalizer-nbands", name)
	if err != nil {
		return nil, err
	}
	return &EqualizerNbands{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// num-bands (uint)
//
// number of different bands to use

func (e *EqualizerNbands) GetNumBands() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-bands")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *EqualizerNbands) SetNumBands(value uint) error {
	return e.Element.SetProperty("num-bands", value)
}

