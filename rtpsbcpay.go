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

type Rtpsbcpay struct {
	Element *gst.Element
}

func NewRtpsbcpay() (*Rtpsbcpay, error) {
	e, err := gst.NewElement("rtpsbcpay")
	if err != nil {
		return nil, err
	}
	return &Rtpsbcpay{Element: e}, nil
}

func NewRtpsbcpayWithName(name string) (*Rtpsbcpay, error) {
	e, err := gst.NewElementWithName("rtpsbcpay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpsbcpay{Element: e}, nil
}

// ----- Properties -----

// min-frames (int)
//
// Minimum quantity of frames to send in one packet (-1 for maximum allowed by the mtu)

func (e *Rtpsbcpay) GetMinFrames() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpsbcpay) SetMinFrames(value int) error {
	return e.Element.SetProperty("min-frames", value)
}

