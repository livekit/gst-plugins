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

type Lamemp3Enc struct {
	Element *gst.Element
}

func NewLamemp3Enc() (*Lamemp3Enc, error) {
	e, err := gst.NewElement("lamemp3enc")
	if err != nil {
		return nil, err
	}
	return &Lamemp3Enc{Element: e}, nil
}

func NewLamemp3EncWithName(name string) (*Lamemp3Enc, error) {
	e, err := gst.NewElementWithName("lamemp3enc", name)
	if err != nil {
		return nil, err
	}
	return &Lamemp3Enc{Element: e}, nil
}

// ----- Properties -----

// bitrate (int)
//
// Bitrate in kbit/sec (Only valid if target is bitrate, for CBR one of 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256 or 320)

func (e *Lamemp3Enc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Lamemp3Enc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// cbr (bool)
//
// Enforce constant bitrate encoding (Only valid if target is bitrate)

func (e *Lamemp3Enc) GetCbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Lamemp3Enc) SetCbr(value bool) error {
	return e.Element.SetProperty("cbr", value)
}

// encoding-engine-quality (GstLameMP3EncEncodingEngineQuality)
//
// Quality/speed of the encoding engine, this does not affect the bitrate!

func (e *Lamemp3Enc) GetEncodingEngineQuality() (interface{}, error) {
	return e.Element.GetProperty("encoding-engine-quality")
}

func (e *Lamemp3Enc) SetEncodingEngineQuality(value interface{}) error {
	return e.Element.SetProperty("encoding-engine-quality", value)
}

// mono (bool)
//
// Enforce mono encoding

func (e *Lamemp3Enc) GetMono() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mono")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Lamemp3Enc) SetMono(value bool) error {
	return e.Element.SetProperty("mono", value)
}

// quality (float32)
//
// VBR Quality from 0 to 10, 0 being the best (Only valid if target is quality)

func (e *Lamemp3Enc) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Lamemp3Enc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

// target (GstLameMP3EncTarget)
//
// Optimize for quality or bitrate

func (e *Lamemp3Enc) GetTarget() (interface{}, error) {
	return e.Element.GetProperty("target")
}

func (e *Lamemp3Enc) SetTarget(value interface{}) error {
	return e.Element.SetProperty("target", value)
}

// ----- Constants -----

type GstLameMp3EncEncodingEngineQuality string

const (
	GstLameMp3EncEncodingEngineQualityFast GstLameMp3EncEncodingEngineQuality = "fast" // fast (0) – Fast
	GstLameMp3EncEncodingEngineQualityStandard GstLameMp3EncEncodingEngineQuality = "standard" // standard (1) – Standard
	GstLameMp3EncEncodingEngineQualityHigh GstLameMp3EncEncodingEngineQuality = "high" // high (2) – High
)

type GstLameMp3EncTarget string

const (
	GstLameMp3EncTargetQuality GstLameMp3EncTarget = "quality" // quality (0) – Quality
	GstLameMp3EncTargetBitrate GstLameMp3EncTarget = "bitrate" // bitrate (1) – Bitrate
)

