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

type Rtpmux struct {
	Element *gst.Element
}

func NewRtpmux() (*Rtpmux, error) {
	e, err := gst.NewElement("rtpmux")
	if err != nil {
		return nil, err
	}
	return &Rtpmux{Element: e}, nil
}

func NewRtpmuxWithName(name string) (*Rtpmux, error) {
	e, err := gst.NewElementWithName("rtpmux", name)
	if err != nil {
		return nil, err
	}
	return &Rtpmux{Element: e}, nil
}

// ----- Properties -----

// seqnum (uint)
//
// The RTP sequence number of the last processed packet

func (e *Rtpmux) GetSeqnum() (uint, error) {
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

func (e *Rtpmux) GetSeqnumOffset() (int, error) {
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

func (e *Rtpmux) SetSeqnumOffset(value int) error {
	return e.Element.SetProperty("seqnum-offset", value)
}

// ssrc (uint)
//
// The SSRC of the packets (default == random)

func (e *Rtpmux) GetSsrc() (uint, error) {
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

func (e *Rtpmux) SetSsrc(value uint) error {
	return e.Element.SetProperty("ssrc", value)
}

// timestamp-offset (int)
//
// Offset to add to all outgoing timestamps (-1 = random)

func (e *Rtpmux) GetTimestampOffset() (int, error) {
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

func (e *Rtpmux) SetTimestampOffset(value int) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

