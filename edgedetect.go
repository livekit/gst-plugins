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

type Edgedetect struct {
	*base.GstBaseTransform
}

func NewEdgedetect() (*Edgedetect, error) {
	e, err := gst.NewElement("edgedetect")
	if err != nil {
		return nil, err
	}
	return &Edgedetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewEdgedetectWithName(name string) (*Edgedetect, error) {
	e, err := gst.NewElementWithName("edgedetect", name)
	if err != nil {
		return nil, err
	}
	return &Edgedetect{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// aperture (int)
//
// Aperture size for Sobel operator (Must be either 3, 5 or 7

func (e *Edgedetect) GetAperture() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("aperture")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Edgedetect) SetAperture(value int) error {
	return e.Element.SetProperty("aperture", value)
}

// mask (bool)
//
// Sets whether the detected edges should be used as a mask on the original input or not

func (e *Edgedetect) GetMask() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Edgedetect) SetMask(value bool) error {
	return e.Element.SetProperty("mask", value)
}

// threshold1 (int)
//
// Threshold value for canny edge detection

func (e *Edgedetect) GetThreshold1() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold1")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Edgedetect) SetThreshold1(value int) error {
	return e.Element.SetProperty("threshold1", value)
}

// threshold2 (int)
//
// Second threshold value for canny edge detection

func (e *Edgedetect) GetThreshold2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Edgedetect) SetThreshold2(value int) error {
	return e.Element.SetProperty("threshold2", value)
}

