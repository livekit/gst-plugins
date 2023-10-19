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

type Chromahold struct {
	*base.GstBaseTransform
}

func NewChromahold() (*Chromahold, error) {
	e, err := gst.NewElement("chromahold")
	if err != nil {
		return nil, err
	}
	return &Chromahold{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewChromaholdWithName(name string) (*Chromahold, error) {
	e, err := gst.NewElementWithName("chromahold", name)
	if err != nil {
		return nil, err
	}
	return &Chromahold{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// target-b (uint)
//
// The Blue target

func (e *Chromahold) GetTargetB() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-b")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Chromahold) SetTargetB(value uint) error {
	return e.Element.SetProperty("target-b", value)
}

// target-g (uint)
//
// The Green target

func (e *Chromahold) GetTargetG() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-g")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Chromahold) SetTargetG(value uint) error {
	return e.Element.SetProperty("target-g", value)
}

// target-r (uint)
//
// The Red target

func (e *Chromahold) GetTargetR() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Chromahold) SetTargetR(value uint) error {
	return e.Element.SetProperty("target-r", value)
}

// tolerance (uint)
//
// Tolerance for the target color

func (e *Chromahold) GetTolerance() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Chromahold) SetTolerance(value uint) error {
	return e.Element.SetProperty("tolerance", value)
}

