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

type Dvdec struct {
	Element *gst.Element
}

func NewDvdec() (*Dvdec, error) {
	e, err := gst.NewElement("dvdec")
	if err != nil {
		return nil, err
	}
	return &Dvdec{Element: e}, nil
}

func NewDvdecWithName(name string) (*Dvdec, error) {
	e, err := gst.NewElementWithName("dvdec", name)
	if err != nil {
		return nil, err
	}
	return &Dvdec{Element: e}, nil
}

// ----- Properties -----

// clamp-chroma (bool)
//
// Clamp chroma

func (e *Dvdec) GetClampChroma() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clamp-chroma")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dvdec) SetClampChroma(value bool) error {
	return e.Element.SetProperty("clamp-chroma", value)
}

// clamp-luma (bool)
//
// Clamp luma

func (e *Dvdec) GetClampLuma() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clamp-luma")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dvdec) SetClampLuma(value bool) error {
	return e.Element.SetProperty("clamp-luma", value)
}

// drop-factor (int)
//
// Only decode Nth frame

func (e *Dvdec) GetDropFactor() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("drop-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dvdec) SetDropFactor(value int) error {
	return e.Element.SetProperty("drop-factor", value)
}

// quality (GstDvdecQualityEnum)
//
// Decoding quality

func (e *Dvdec) GetQuality() (GstDvdecQualityEnum, error) {
	var value GstDvdecQualityEnum
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvdecQualityEnum)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvdecQualityEnum")
	}
	return value, nil
}

func (e *Dvdec) SetQuality(value GstDvdecQualityEnum) error {
	e.Element.SetArg("quality", string(value))
	return nil
}

// ----- Constants -----

type GstDvdecQualityEnum string

const (
	GstDvdecQualityEnumFastest GstDvdecQualityEnum = "fastest" // fastest (0) – Monochrome, DC (Fastest)
	GstDvdecQualityEnumMonochromeAc GstDvdecQualityEnum = "monochrome-ac" // monochrome-ac (1) – Monochrome, first AC coefficient
	GstDvdecQualityEnumMonochromeBest GstDvdecQualityEnum = "monochrome-best" // monochrome-best (2) – Monochrome, highest quality
	GstDvdecQualityEnumColourFastest GstDvdecQualityEnum = "colour-fastest" // colour-fastest (3) – Colour, DC, fastest
	GstDvdecQualityEnumColourAc GstDvdecQualityEnum = "colour-ac" // colour-ac (4) – Colour, using only the first AC coefficient
	GstDvdecQualityEnumBest GstDvdecQualityEnum = "best" // best (5) – Highest quality colour decoding
)

