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

type Faac struct {
	Element *gst.Element
}

func NewFaac() (*Faac, error) {
	e, err := gst.NewElement("faac")
	if err != nil {
		return nil, err
	}
	return &Faac{Element: e}, nil
}

func NewFaacWithName(name string) (*Faac, error) {
	e, err := gst.NewElementWithName("faac", name)
	if err != nil {
		return nil, err
	}
	return &Faac{Element: e}, nil
}

// ----- Properties -----

// bitrate (int)
//
// Average Bitrate (ABR) in bits/sec

func (e *Faac) GetBitrate() (int, error) {
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

func (e *Faac) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// midside (bool)
//
// Allow mid/side encoding

func (e *Faac) GetMidside() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("midside")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Faac) SetMidside(value bool) error {
	return e.Element.SetProperty("midside", value)
}

// quality (int)
//
// Variable bitrate (VBR) quantizer quality in %%

func (e *Faac) GetQuality() (int, error) {
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

func (e *Faac) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

// rate-control (GstFaacBrtype)
//
// Encoding bitrate type (VBR/ABR)

func (e *Faac) GetRateControl() (GstFaacBrtype, error) {
	var value GstFaacBrtype
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFaacBrtype)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFaacBrtype")
	}
	return value, nil
}

func (e *Faac) SetRateControl(value GstFaacBrtype) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

// shortctl (GstFaacShortCtl)
//
// Block type encorcing

func (e *Faac) GetShortctl() (GstFaacShortCtl, error) {
	var value GstFaacShortCtl
	var ok bool
	v, err := e.Element.GetProperty("shortctl")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFaacShortCtl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFaacShortCtl")
	}
	return value, nil
}

func (e *Faac) SetShortctl(value GstFaacShortCtl) error {
	e.Element.SetArg("shortctl", string(value))
	return nil
}

// tns (bool)
//
// Use temporal noise shaping

func (e *Faac) GetTns() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("tns")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Faac) SetTns(value bool) error {
	return e.Element.SetProperty("tns", value)
}

// ----- Constants -----

type GstFaacBrtype string

const (
	GstFaacBrtypeVbrEncoding GstFaacBrtype = "VBR encoding" // VBR encoding (1) – VBR
	GstFaacBrtypeAbrEncoding GstFaacBrtype = "ABR encoding" // ABR encoding (2) – ABR
)

type GstFaacShortCtl string

const (
	GstFaacShortCtlNormalBlockType GstFaacShortCtl = "Normal block type" // Normal block type (0) – SHORTCTL_NORMAL
	GstFaacShortCtlNoShortBlocks GstFaacShortCtl = "No short blocks" // No short blocks (1) – SHORTCTL_NOSHORT
	GstFaacShortCtlNoLongBlocks GstFaacShortCtl = "No long blocks" // No long blocks (2) – SHORTCTL_NOLONG
)

