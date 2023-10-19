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
	"github.com/tinyzimmer/go-gst/gst"
)

type Rtpilbcdepay struct {
	Element *gst.Element
}

func NewRtpilbcdepay() (*Rtpilbcdepay, error) {
	e, err := gst.NewElement("rtpilbcdepay")
	if err != nil {
		return nil, err
	}
	return &Rtpilbcdepay{Element: e}, nil
}

func NewRtpilbcdepayWithName(name string) (*Rtpilbcdepay, error) {
	e, err := gst.NewElementWithName("rtpilbcdepay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpilbcdepay{Element: e}, nil
}

// ----- Properties -----

// mode (GstILbcmode)
//
// iLBC frame mode

func (e *Rtpilbcdepay) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *Rtpilbcdepay) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ----- Constants -----

type ILbcmode string

const (
	ILbcmode20Ms ILbcmode = "20ms" // 20ms (20) – 20ms frames
	ILbcmode30Ms ILbcmode = "30ms" // 30ms (30) – 30ms frames
)

