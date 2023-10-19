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

type Rtprtxqueue struct {
	Element *gst.Element
}

func NewRtprtxqueue() (*Rtprtxqueue, error) {
	e, err := gst.NewElement("rtprtxqueue")
	if err != nil {
		return nil, err
	}
	return &Rtprtxqueue{Element: e}, nil
}

func NewRtprtxqueueWithName(name string) (*Rtprtxqueue, error) {
	e, err := gst.NewElementWithName("rtprtxqueue", name)
	if err != nil {
		return nil, err
	}
	return &Rtprtxqueue{Element: e}, nil
}

// ----- Properties -----

// fulfilled-requests (uint)
//
// Number of fulfilled retransmission requests

func (e *Rtprtxqueue) GetFulfilledRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("fulfilled-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// max-size-packets (uint)
//
// Amount of packets to queue (0 = unlimited)

func (e *Rtprtxqueue) GetMaxSizePackets() (uint, error) {
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

func (e *Rtprtxqueue) SetMaxSizePackets(value uint) error {
	return e.Element.SetProperty("max-size-packets", value)
}

// max-size-time (uint)
//
// Amount of ms to queue (0 = unlimited)

func (e *Rtprtxqueue) GetMaxSizeTime() (uint, error) {
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

func (e *Rtprtxqueue) SetMaxSizeTime(value uint) error {
	return e.Element.SetProperty("max-size-time", value)
}

// requests (uint)
//
// Total number of retransmission requests

func (e *Rtprtxqueue) GetRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

