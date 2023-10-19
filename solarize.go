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

type Solarize struct {
	*base.GstBaseTransform
}

func NewSolarize() (*Solarize, error) {
	e, err := gst.NewElement("solarize")
	if err != nil {
		return nil, err
	}
	return &Solarize{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSolarizeWithName(name string) (*Solarize, error) {
	e, err := gst.NewElementWithName("solarize", name)
	if err != nil {
		return nil, err
	}
	return &Solarize{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// end (uint)
//
// End parameter

func (e *Solarize) GetEnd() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("end")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Solarize) SetEnd(value uint) error {
	return e.Element.SetProperty("end", value)
}

// start (uint)
//
// Start parameter

func (e *Solarize) GetStart() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("start")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Solarize) SetStart(value uint) error {
	return e.Element.SetProperty("start", value)
}

// threshold (uint)
//
// Threshold parameter

func (e *Solarize) GetThreshold() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Solarize) SetThreshold(value uint) error {
	return e.Element.SetProperty("threshold", value)
}

