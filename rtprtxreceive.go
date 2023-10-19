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

type Rtprtxreceive struct {
	Element *gst.Element
}

func NewRtprtxreceive() (*Rtprtxreceive, error) {
	e, err := gst.NewElement("rtprtxreceive")
	if err != nil {
		return nil, err
	}
	return &Rtprtxreceive{Element: e}, nil
}

func NewRtprtxreceiveWithName(name string) (*Rtprtxreceive, error) {
	e, err := gst.NewElementWithName("rtprtxreceive", name)
	if err != nil {
		return nil, err
	}
	return &Rtprtxreceive{Element: e}, nil
}

// ----- Properties -----

// num-rtx-assoc-packets (uint)
//
// Number of retransmission packets correctly associated with retransmission requests

func (e *Rtprtxreceive) GetNumRtxAssocPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-assoc-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// num-rtx-packets (uint)
//
// Number of retransmission packets received

func (e *Rtprtxreceive) GetNumRtxPackets() (uint, error) {
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

func (e *Rtprtxreceive) GetNumRtxRequests() (uint, error) {
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

func (e *Rtprtxreceive) GetPayloadTypeMap() (interface{}, error) {
	return e.Element.GetProperty("payload-type-map")
}

func (e *Rtprtxreceive) SetPayloadTypeMap(value interface{}) error {
	return e.Element.SetProperty("payload-type-map", value)
}

// ssrc-map (GstStructure)
//
// Map of SSRCs to their retransmission SSRCs for SSRC-multiplexed mode.

