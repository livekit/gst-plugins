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

type Equalizer10Bands struct {
	*base.GstBaseTransform
}

func NewEqualizer10Bands() (*Equalizer10Bands, error) {
	e, err := gst.NewElement("equalizer-10bands")
	if err != nil {
		return nil, err
	}
	return &Equalizer10Bands{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewEqualizer10BandsWithName(name string) (*Equalizer10Bands, error) {
	e, err := gst.NewElementWithName("equalizer-10bands", name)
	if err != nil {
		return nil, err
	}
	return &Equalizer10Bands{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// band0 (float64)
//
// gain for the frequency band 29 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand0() (float64, error) {
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

func (e *Equalizer10Bands) SetBand0(value float64) error {
	return e.Element.SetProperty("band0", value)
}

// band1 (float64)
//
// gain for the frequency band 59 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand1() (float64, error) {
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

func (e *Equalizer10Bands) SetBand1(value float64) error {
	return e.Element.SetProperty("band1", value)
}

// band2 (float64)
//
// gain for the frequency band 119 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand2() (float64, error) {
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

func (e *Equalizer10Bands) SetBand2(value float64) error {
	return e.Element.SetProperty("band2", value)
}

// band3 (float64)
//
// gain for the frequency band 237 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand3() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand3(value float64) error {
	return e.Element.SetProperty("band3", value)
}

// band4 (float64)
//
// gain for the frequency band 474 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand4() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand4(value float64) error {
	return e.Element.SetProperty("band4", value)
}

// band5 (float64)
//
// gain for the frequency band 947 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand5() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand5(value float64) error {
	return e.Element.SetProperty("band5", value)
}

// band6 (float64)
//
// gain for the frequency band 1889 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand6() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand6(value float64) error {
	return e.Element.SetProperty("band6", value)
}

// band7 (float64)
//
// gain for the frequency band 3770 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand7() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand7(value float64) error {
	return e.Element.SetProperty("band7", value)
}

// band8 (float64)
//
// gain for the frequency band 7523 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand8() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand8(value float64) error {
	return e.Element.SetProperty("band8", value)
}

// band9 (float64)
//
// gain for the frequency band 15011 Hz, ranging from -24 dB to +12 dB

func (e *Equalizer10Bands) GetBand9() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("band9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Equalizer10Bands) SetBand9(value float64) error {
	return e.Element.SetProperty("band9", value)
}

