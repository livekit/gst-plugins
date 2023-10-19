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

type Audiochannelmix struct {
	*base.GstBaseTransform
}

func NewAudiochannelmix() (*Audiochannelmix, error) {
	e, err := gst.NewElement("audiochannelmix")
	if err != nil {
		return nil, err
	}
	return &Audiochannelmix{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiochannelmixWithName(name string) (*Audiochannelmix, error) {
	e, err := gst.NewElementWithName("audiochannelmix", name)
	if err != nil {
		return nil, err
	}
	return &Audiochannelmix{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// left-to-left (float64)
//
// Left channel to left channel gain

func (e *Audiochannelmix) GetLeftToLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("left-to-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Audiochannelmix) SetLeftToLeft(value float64) error {
	return e.Element.SetProperty("left-to-left", value)
}

// left-to-right (float64)
//
// Left channel to right channel gain

func (e *Audiochannelmix) GetLeftToRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("left-to-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Audiochannelmix) SetLeftToRight(value float64) error {
	return e.Element.SetProperty("left-to-right", value)
}

// right-to-left (float64)
//
// Right channel to left channel gain

func (e *Audiochannelmix) GetRightToLeft() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("right-to-left")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Audiochannelmix) SetRightToLeft(value float64) error {
	return e.Element.SetProperty("right-to-left", value)
}

// right-to-right (float64)
//
// Right channel to right channel gain

func (e *Audiochannelmix) GetRightToRight() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("right-to-right")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Audiochannelmix) SetRightToRight(value float64) error {
	return e.Element.SetProperty("right-to-right", value)
}

