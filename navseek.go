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

type Navseek struct {
	*base.GstBaseTransform
}

func NewNavseek() (*Navseek, error) {
	e, err := gst.NewElement("navseek")
	if err != nil {
		return nil, err
	}
	return &Navseek{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewNavseekWithName(name string) (*Navseek, error) {
	e, err := gst.NewElementWithName("navseek", name)
	if err != nil {
		return nil, err
	}
	return &Navseek{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// hold-eos (bool)
//
// Hold eos until the next 'Return' keystroke.

func (e *Navseek) GetHoldEos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hold-eos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Navseek) SetHoldEos(value bool) error {
	return e.Element.SetProperty("hold-eos", value)
}

// seek-offset (float64)
//
// Time in seconds to seek by

func (e *Navseek) GetSeekOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("seek-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Navseek) SetSeekOffset(value float64) error {
	return e.Element.SetProperty("seek-offset", value)
}

