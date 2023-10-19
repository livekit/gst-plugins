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

type Rtponviftimestamp struct {
	Element *gst.Element
}

func NewRtponviftimestamp() (*Rtponviftimestamp, error) {
	e, err := gst.NewElement("rtponviftimestamp")
	if err != nil {
		return nil, err
	}
	return &Rtponviftimestamp{Element: e}, nil
}

func NewRtponviftimestampWithName(name string) (*Rtponviftimestamp, error) {
	e, err := gst.NewElementWithName("rtponviftimestamp", name)
	if err != nil {
		return nil, err
	}
	return &Rtponviftimestamp{Element: e}, nil
}

// ----- Properties -----

// cseq (uint)
//
// The RTSP CSeq which initiated the playback

func (e *Rtponviftimestamp) GetCseq() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cseq")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtponviftimestamp) SetCseq(value uint) error {
	return e.Element.SetProperty("cseq", value)
}

// drop-out-of-segment (bool)
//
// Whether the element should drop buffers that fall outside the segment, not part of the specification but allows full reverse playback.

func (e *Rtponviftimestamp) GetDropOutOfSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-out-of-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtponviftimestamp) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

// ntp-offset (uint64)
//
// Offset between the pipeline running time and the absolute UTC time, in nano-seconds since 1900 (-1 for automatic computation)

func (e *Rtponviftimestamp) GetNtpOffset() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ntp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtponviftimestamp) SetNtpOffset(value uint64) error {
	return e.Element.SetProperty("ntp-offset", value)
}

// set-e-bit (bool)
//
// If the element should set the 'E' bit as defined in the ONVIF RTP extension. This increases latency by one packet

func (e *Rtponviftimestamp) GetSetEBit() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("set-e-bit")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtponviftimestamp) SetSetEBit(value bool) error {
	return e.Element.SetProperty("set-e-bit", value)
}

// set-t-bit (bool)
//
// If the element should set the 'T' bit as defined in the ONVIF RTP extension. This increases latency by one packet

func (e *Rtponviftimestamp) GetSetTBit() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("set-t-bit")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtponviftimestamp) SetSetTBit(value bool) error {
	return e.Element.SetProperty("set-t-bit", value)
}

// use-reference-timestamps (bool)
//
// Whether to obtain timestamps from reference timestamp meta instead of using
// the ntp-offset method. If enabled then timestamps are expected to be
// attached to the buffers, and in that case ntp-offset should not be
// configured.

