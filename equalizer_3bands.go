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

type Equalizer3Bands struct {
	*base.GstBaseTransform
}

func NewEqualizer3Bands() (*Equalizer3Bands, error) {
	e, err := gst.NewElement("equalizer-3bands")
	if err != nil {
		return nil, err
	}
	return &Equalizer3Bands{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewEqualizer3BandsWithName(name string) (*Equalizer3Bands, error) {
	e, err := gst.NewElementWithName("equalizer-3bands", name)
	if err != nil {
		return nil, err
	}
	return &Equalizer3Bands{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// band0 (float64)
//
// gain for the frequency band 100 Hz, ranging from -24.0 to +12.0

func (e *Equalizer3Bands) GetBand0() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer3Bands) SetBand0(value float64) error {
	return e.Element.SetProperty("band0", value)
}

// band1 (float64)
//
// gain for the frequency band 1100 Hz, ranging from -24.0 to +12.0

func (e *Equalizer3Bands) GetBand1() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer3Bands) SetBand1(value float64) error {
	return e.Element.SetProperty("band1", value)
}

// band2 (float64)
//
// gain for the frequency band 11 kHz, ranging from -24.0 to +12.0

func (e *Equalizer3Bands) GetBand2() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer3Bands) SetBand2(value float64) error {
	return e.Element.SetProperty("band2", value)
}

