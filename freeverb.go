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

type Freeverb struct {
	*base.GstBaseTransform
}

func NewFreeverb() (*Freeverb, error) {
	e, err := gst.NewElement("freeverb")
	if err != nil {
		return nil, err
	}
	return &Freeverb{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFreeverbWithName(name string) (*Freeverb, error) {
	e, err := gst.NewElementWithName("freeverb", name)
	if err != nil {
		return nil, err
	}
	return &Freeverb{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// damping (float32)
//
// Damping of high frequencies

func (e *Freeverb) GetDamping() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("damping")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Freeverb) SetDamping(value float32) error {
	return e.Element.SetProperty("damping", value)
}

// level (float32)
//
// dry/wet level

func (e *Freeverb) GetLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Freeverb) SetLevel(value float32) error {
	return e.Element.SetProperty("level", value)
}

// room-size (float32)
//
// Size of the simulated room

func (e *Freeverb) GetRoomSize() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("room-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Freeverb) SetRoomSize(value float32) error {
	return e.Element.SetProperty("room-size", value)
}

// width (float32)
//
// Stereo panorama width

func (e *Freeverb) GetWidth() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Freeverb) SetWidth(value float32) error {
	return e.Element.SetProperty("width", value)
}

