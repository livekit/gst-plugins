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

type Rtprtxsend struct {
	Element *gst.Element
}

func NewRtprtxsend() (*Rtprtxsend, error) {
	e, err := gst.NewElement("rtprtxsend")
	if err != nil {
		return nil, err
	}
	return &Rtprtxsend{Element: e}, nil
}

func NewRtprtxsendWithName(name string) (*Rtprtxsend, error) {
	e, err := gst.NewElementWithName("rtprtxsend", name)
	if err != nil {
		return nil, err
	}
	return &Rtprtxsend{Element: e}, nil
}

// ----- Properties -----

// clock-rate-map (GstStructure)
//
// Map of payload types to their clock rates

func (e *Rtprtxsend) GetClockRateMap() (interface{}, error) {
	return e.Element.GetProperty("clock-rate-map")
}

func (e *Rtprtxsend) SetClockRateMap(value interface{}) error {
	return e.Element.SetProperty("clock-rate-map", value)
}

// max-size-packets (uint)
//
// Amount of packets to queue (0 = unlimited)

func (e *Rtprtxsend) GetMaxSizePackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtprtxsend) SetMaxSizePackets(value uint) error {
	return e.Element.SetProperty("max-size-packets", value)
}

// max-size-time (uint)
//
// Amount of ms to queue (0 = unlimited)

func (e *Rtprtxsend) GetMaxSizeTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtprtxsend) SetMaxSizeTime(value uint) error {
	return e.Element.SetProperty("max-size-time", value)
}

// num-rtx-packets (uint)
//
// Number of retransmission packets sent

func (e *Rtprtxsend) GetNumRtxPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// num-rtx-requests (uint)
//
// Number of retransmission events received

func (e *Rtprtxsend) GetNumRtxRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// payload-type-map (GstStructure)
//
// Map of original payload types to their retransmission payload types

func (e *Rtprtxsend) GetPayloadTypeMap() (interface{}, error) {
	return e.Element.GetProperty("payload-type-map")
}

func (e *Rtprtxsend) SetPayloadTypeMap(value interface{}) error {
	return e.Element.SetProperty("payload-type-map", value)
}

// ssrc-map (GstStructure)
//
// Map of SSRCs to their retransmission SSRCs for SSRC-multiplexed mode (default = random)

func (e *Rtprtxsend) GetSsrcMap() (interface{}, error) {
	return e.Element.GetProperty("ssrc-map")
}

func (e *Rtprtxsend) SetSsrcMap(value interface{}) error {
	return e.Element.SetProperty("ssrc-map", value)
}

