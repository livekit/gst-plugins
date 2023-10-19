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

type Handdetect struct {
	*base.GstBaseTransform
}

func NewHanddetect() (*Handdetect, error) {
	e, err := gst.NewElement("handdetect")
	if err != nil {
		return nil, err
	}
	return &Handdetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewHanddetectWithName(name string) (*Handdetect, error) {
	e, err := gst.NewElementWithName("handdetect", name)
	if err != nil {
		return nil, err
	}
	return &Handdetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// ROI-HEIGHT (int)
//
// HEIGHT of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages

func (e *Handdetect) GetRoiHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-HEIGHT")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Handdetect) SetRoiHeight(value int) error {
	return e.Element.SetProperty("ROI-HEIGHT", value)
}

// ROI-WIDTH (int)
//
// WIDTH of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages

func (e *Handdetect) GetRoiWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-WIDTH")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Handdetect) SetRoiWidth(value int) error {
	return e.Element.SetProperty("ROI-WIDTH", value)
}

// ROI-X (int)
//
// X of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages

func (e *Handdetect) GetRoiX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-X")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Handdetect) SetRoiX(value int) error {
	return e.Element.SetProperty("ROI-X", value)
}

// ROI-Y (int)
//
// Y of left-top pointer in region of interest
// Gestures in the defined region of interest will emit messages

func (e *Handdetect) GetRoiY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ROI-Y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Handdetect) SetRoiY(value int) error {
	return e.Element.SetProperty("ROI-Y", value)
}

// display (bool)
//
// Whether the detected hands are highlighted in output frame

func (e *Handdetect) GetDisplay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("display")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Handdetect) SetDisplay(value bool) error {
	return e.Element.SetProperty("display", value)
}

// profile-fist (string)
//
// Location of HAAR cascade file (fist gesture)

func (e *Handdetect) GetProfileFist() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile-fist")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Handdetect) SetProfileFist(value string) error {
	return e.Element.SetProperty("profile-fist", value)
}

// profile-palm (string)
//
// Location of HAAR cascade file (palm gesture)

func (e *Handdetect) GetProfilePalm() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("profile-palm")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Handdetect) SetProfilePalm(value string) error {
	return e.Element.SetProperty("profile-palm", value)
}

