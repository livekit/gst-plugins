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

type Rtpsink struct {
	Element *gst.Element
}

func NewRtpsink() (*Rtpsink, error) {
	e, err := gst.NewElement("rtpsink")
	if err != nil {
		return nil, err
	}
	return &Rtpsink{Element: e}, nil
}

func NewRtpsinkWithName(name string) (*Rtpsink, error) {
	e, err := gst.NewElementWithName("rtpsink", name)
	if err != nil {
		return nil, err
	}
	return &Rtpsink{Element: e}, nil
}

// ----- Properties -----

// address (string)
//
// Address to receive packets from (can be IPv4 or IPv6).

func (e *Rtpsink) GetAddress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("address")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtpsink) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

// multicast-iface (string)
//
// The networkinterface on which to join the multicast group

func (e *Rtpsink) GetMulticastIface() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("multicast-iface")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtpsink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// port (uint)
//
// The port to listen to RTP packets, the RTCP port is this value
// +1. This port must be an even number.

func (e *Rtpsink) GetPort() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("port")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpsink) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

// ttl (int)
//
// Set the unicast TTL parameter.

func (e *Rtpsink) GetTtl() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ttl")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpsink) SetTtl(value int) error {
	return e.Element.SetProperty("ttl", value)
}

// ttl-mc (int)
//
// Set the multicast TTL parameter.

func (e *Rtpsink) GetTtlMc() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ttl-mc")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpsink) SetTtlMc(value int) error {
	return e.Element.SetProperty("ttl-mc", value)
}

// uri (string)
//
// uri to stream RTP to. All GStreamer parameters can be
// encoded in the URI, this URI format is RFC compliant.

func (e *Rtpsink) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtpsink) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

