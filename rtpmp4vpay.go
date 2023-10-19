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

type Rtpmp4Vpay struct {
	Element *gst.Element
}

func NewRtpmp4Vpay() (*Rtpmp4Vpay, error) {
	e, err := gst.NewElement("rtpmp4vpay")
	if err != nil {
		return nil, err
	}
	return &Rtpmp4Vpay{Element: e}, nil
}

func NewRtpmp4VpayWithName(name string) (*Rtpmp4Vpay, error) {
	e, err := gst.NewElementWithName("rtpmp4vpay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpmp4Vpay{Element: e}, nil
}

// ----- Properties -----

// config-interval (int)
//
// Send Config Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)

func (e *Rtpmp4Vpay) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpmp4Vpay) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

