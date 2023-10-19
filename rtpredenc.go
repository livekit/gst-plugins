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

type Rtpredenc struct {
	Element *gst.Element
}

func NewRtpredenc() (*Rtpredenc, error) {
	e, err := gst.NewElement("rtpredenc")
	if err != nil {
		return nil, err
	}
	return &Rtpredenc{Element: e}, nil
}

func NewRtpredencWithName(name string) (*Rtpredenc, error) {
	e, err := gst.NewElementWithName("rtpredenc", name)
	if err != nil {
		return nil, err
	}
	return &Rtpredenc{Element: e}, nil
}

// ----- Properties -----

// allow-no-red-blocks (bool)
//
// true - can produce RED packets even without redundant blocks (distance==0) false - RED packets will be produced only if distance>0

func (e *Rtpredenc) GetAllowNoRedBlocks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-no-red-blocks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpredenc) SetAllowNoRedBlocks(value bool) error {
	return e.Element.SetProperty("allow-no-red-blocks", value)
}

// distance (uint)
//
// Tells which media packet to use as a redundant block (0 - no redundant blocks, 1 to use previous packet, 2 to use the packet before previous, etc.)

func (e *Rtpredenc) GetDistance() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpredenc) SetDistance(value uint) error {
	return e.Element.SetProperty("distance", value)
}

// pt (int)
//
// Payload type FEC packets (-1 disable)

func (e *Rtpredenc) GetPt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpredenc) SetPt(value int) error {
	return e.Element.SetProperty("pt", value)
}

// sent (uint)
//
// Count of sent packets

func (e *Rtpredenc) GetSent() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sent")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

