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

type Rtpdtmfdepay struct {
	Element *gst.Element
}

func NewRtpdtmfdepay() (*Rtpdtmfdepay, error) {
	e, err := gst.NewElement("rtpdtmfdepay")
	if err != nil {
		return nil, err
	}
	return &Rtpdtmfdepay{Element: e}, nil
}

func NewRtpdtmfdepayWithName(name string) (*Rtpdtmfdepay, error) {
	e, err := gst.NewElementWithName("rtpdtmfdepay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpdtmfdepay{Element: e}, nil
}

// ----- Properties -----

// max-duration (uint)
//
// The maxumimum duration (ms) of the outgoing soundpacket. (0 = no limit)

func (e *Rtpdtmfdepay) GetMaxDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpdtmfdepay) SetMaxDuration(value uint) error {
	return e.Element.SetProperty("max-duration", value)
}

// unit-time (uint)
//
// The smallest unit (ms) the duration must be a multiple of (0 disables it)

func (e *Rtpdtmfdepay) GetUnitTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("unit-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpdtmfdepay) SetUnitTime(value uint) error {
	return e.Element.SetProperty("unit-time", value)
}

