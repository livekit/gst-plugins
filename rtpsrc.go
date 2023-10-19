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

type Rtpsrc struct {
	Element *gst.Element
}

func NewRtpsrc() (*Rtpsrc, error) {
	e, err := gst.NewElement("rtpsrc")
	if err != nil {
		return nil, err
	}
	return &Rtpsrc{Element: e}, nil
}

func NewRtpsrcWithName(name string) (*Rtpsrc, error) {
	e, err := gst.NewElementWithName("rtpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Rtpsrc{Element: e}, nil
}

// ----- Properties -----

// address (string)
//
// Address to receive packets from (can be IPv4 or IPv6).

func (e *Rtpsrc) GetAddress() (string, error) {
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

func (e *Rtpsrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

// caps (GstCaps)
//
// The RTP caps of the incoming stream.

func (e *Rtpsrc) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Rtpsrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// encoding-name (string)
//
// Set the encoding name of the stream to use. This is a short-hand for
// the full caps and maps typically to the encoding-name in the RTP caps.

func (e *Rtpsrc) GetEncodingName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("encoding-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtpsrc) SetEncodingName(value string) error {
	return e.Element.SetProperty("encoding-name", value)
}

// latency (uint)
//
// Set the size of the latency buffer in the
// GstRtpBin/GstRtpJitterBuffer to compensate for network jitter.

func (e *Rtpsrc) GetLatency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpsrc) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

// multicast-iface (string)
//
// The network interface on which to join the multicast group.This allows multiple interfaces separated by comma. ("eth0,eth1")

func (e *Rtpsrc) GetMulticastIface() (string, error) {
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

func (e *Rtpsrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// port (uint)
//
// The port to listen to RTP packets, the RTCP port is this value
// +1. This port must be an even number.

func (e *Rtpsrc) GetPort() (uint, error) {
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

func (e *Rtpsrc) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

// ttl (int)
//
// Set the unicast TTL parameter. In RTP this of importance for RTCP.

func (e *Rtpsrc) GetTtl() (int, error) {
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

func (e *Rtpsrc) SetTtl(value int) error {
	return e.Element.SetProperty("ttl", value)
}

// ttl-mc (int)
//
// Set the multicast TTL parameter. In RTP this of importance for RTCP.

func (e *Rtpsrc) GetTtlMc() (int, error) {
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

func (e *Rtpsrc) SetTtlMc(value int) error {
	return e.Element.SetProperty("ttl-mc", value)
}

// uri (string)
//
// uri to an RTP from. All GStreamer parameters can be
// encoded in the URI, this URI format is RFC compliant.

func (e *Rtpsrc) GetUri() (string, error) {
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

func (e *Rtpsrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

