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

type Rtppassthroughpay struct {
	Element *gst.Element
}

func NewRtppassthroughpay() (*Rtppassthroughpay, error) {
	e, err := gst.NewElement("rtppassthroughpay")
	if err != nil {
		return nil, err
	}
	return &Rtppassthroughpay{Element: e}, nil
}

func NewRtppassthroughpayWithName(name string) (*Rtppassthroughpay, error) {
	e, err := gst.NewElementWithName("rtppassthroughpay", name)
	if err != nil {
		return nil, err
	}
	return &Rtppassthroughpay{Element: e}, nil
}

// ----- Properties -----

// mtu (uint)
//
// Setting this property has no effect on this element, it is here and it
// is writable only to emulate a proper RTP payloader.

func (e *Rtppassthroughpay) GetMtu() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("mtu")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtppassthroughpay) SetMtu(value uint) error {
	return e.Element.SetProperty("mtu", value)
}

// pt (uint)
//
// If set this will override the payload of the incoming RTP packets.
// If not set the payload type will be same as incoming RTP packets.

func (e *Rtppassthroughpay) GetPt() (uint, error) {
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

func (e *Rtppassthroughpay) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

// seqnum (uint)

func (e *Rtppassthroughpay) GetSeqnum() (uint, error) {
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

func (e *Rtppassthroughpay) GetSeqnumOffset() (int, error) {
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

func (e *Rtppassthroughpay) SetSeqnumOffset(value int) error {
	return e.Element.SetProperty("seqnum-offset", value)
}

// stats (GstStructure)

func (e *Rtppassthroughpay) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// timestamp (uint)

func (e *Rtppassthroughpay) GetTimestamp() (uint, error) {
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

// timestamp-offset (uint)

func (e *Rtppassthroughpay) GetTimestampOffset() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtppassthroughpay) SetTimestampOffset(value uint) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

