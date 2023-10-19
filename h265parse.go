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

type H265Parse struct {
	*base.GstBaseParse
}

func NewH265Parse() (*H265Parse, error) {
	e, err := gst.NewElement("h265parse")
	if err != nil {
		return nil, err
	}
	return &H265Parse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewH265ParseWithName(name string) (*H265Parse, error) {
	e, err := gst.NewElementWithName("h265parse", name)
	if err != nil {
		return nil, err
	}
	return &H265Parse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// config-interval (int)
//
// Send VPS, SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)

func (e *H265Parse) GetConfigInterval() (int, error) {
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

func (e *H265Parse) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

