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

type Rtpdtmfsrc struct {
	*base.GstBaseSrc
}

func NewRtpdtmfsrc() (*Rtpdtmfsrc, error) {
	e, err := gst.NewElement("rtpdtmfsrc")
	if err != nil {
		return nil, err
	}
	return &Rtpdtmfsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewRtpdtmfsrcWithName(name string) (*Rtpdtmfsrc, error) {
	e, err := gst.NewElementWithName("rtpdtmfsrc", name)
	if err != nil {
		return nil, err
	}
	return &Rtpdtmfsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// clock-rate (uint)
//
// The clock-rate at which to generate the dtmf packets

func (e *Rtpdtmfsrc) GetClockRate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("clock-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpdtmfsrc) SetClockRate(value uint) error {
	return e.Element.SetProperty("clock-rate", value)
}

// packet-redundancy (uint)
//
// Number of packets to send to indicate start and stop dtmf events

func (e *Rtpdtmfsrc) GetPacketRedundancy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("packet-redundancy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpdtmfsrc) SetPacketRedundancy(value uint) error {
	return e.Element.SetProperty("packet-redundancy", value)
}

// pt (uint)
//
// The payload type of the packets

func (e *Rtpdtmfsrc) GetPt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpdtmfsrc) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

// seqnum (uint)
//
// The RTP sequence number of the last processed packet

func (e *Rtpdtmfsrc) GetSeqnum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// seqnum-offset (int)
//
// Offset to add to all outgoing seqnum (-1 = random)

func (e *Rtpdtmfsrc) GetSeqnumOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("seqnum-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpdtmfsrc) SetSeqnumOffset(value int) error {
	return e.Element.SetProperty("seqnum-offset", value)
}

// ssrc (uint)
//
// The SSRC of the packets (-1 == random)

func (e *Rtpdtmfsrc) GetSsrc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ssrc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpdtmfsrc) SetSsrc(value uint) error {
	return e.Element.SetProperty("ssrc", value)
}

// timestamp (uint)
//
// The RTP timestamp of the last processed packet

func (e *Rtpdtmfsrc) GetTimestamp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// timestamp-offset (int)
//
// Offset to add to all outgoing timestamps (-1 = random)

func (e *Rtpdtmfsrc) GetTimestampOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpdtmfsrc) SetTimestampOffset(value int) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

