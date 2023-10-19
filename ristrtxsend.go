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

type Ristrtxsend struct {
	Element *gst.Element
}

func NewRistrtxsend() (*Ristrtxsend, error) {
	e, err := gst.NewElement("ristrtxsend")
	if err != nil {
		return nil, err
	}
	return &Ristrtxsend{Element: e}, nil
}

func NewRistrtxsendWithName(name string) (*Ristrtxsend, error) {
	e, err := gst.NewElementWithName("ristrtxsend", name)
	if err != nil {
		return nil, err
	}
	return &Ristrtxsend{Element: e}, nil
}

// ----- Properties -----

// max-size-packets (uint)
//
// Amount of packets to queue (0 = unlimited)

func (e *Ristrtxsend) GetMaxSizePackets() (uint, error) {
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

func (e *Ristrtxsend) SetMaxSizePackets(value uint) error {
	return e.Element.SetProperty("max-size-packets", value)
}

// max-size-time (uint)
//
// Amount of ms to queue (0 = unlimited)

func (e *Ristrtxsend) GetMaxSizeTime() (uint, error) {
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

func (e *Ristrtxsend) SetMaxSizeTime(value uint) error {
	return e.Element.SetProperty("max-size-time", value)
}

// num-rtx-packets (uint)
//
// Number of retransmission packets sent

func (e *Ristrtxsend) GetNumRtxPackets() (uint, error) {
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

func (e *Ristrtxsend) GetNumRtxRequests() (uint, error) {
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

