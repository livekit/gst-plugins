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

type Ristsrc struct {
	Element *gst.Element
}

func NewRistsrc() (*Ristsrc, error) {
	e, err := gst.NewElement("ristsrc")
	if err != nil {
		return nil, err
	}
	return &Ristsrc{Element: e}, nil
}

func NewRistsrcWithName(name string) (*Ristsrc, error) {
	e, err := gst.NewElementWithName("ristsrc", name)
	if err != nil {
		return nil, err
	}
	return &Ristsrc{Element: e}, nil
}

// ----- Properties -----

// address (string)
//
// Address to receive packets from (can be IPv4 or IPv6).

func (e *Ristsrc) GetAddress() (string, error) {
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

func (e *Ristsrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

// bonding-addresses (string)
//
// Comma (,) separated list of 

func (e *Ristsrc) GetBondingAddresses() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("bonding-addresses")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Ristsrc) SetBondingAddresses(value string) error {
	return e.Element.SetProperty("bonding-addresses", value)
}

// caps (GstCaps)
//
// The RTP caps of the incoming RIST stream.

func (e *Ristsrc) GetCaps() (*gst.Caps, error) {
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

func (e *Ristsrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// cname (string)
//
// Set the CNAME in the SDES block of the receiver report.

func (e *Ristsrc) GetCname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("cname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Ristsrc) SetCname(value string) error {
	return e.Element.SetProperty("cname", value)
}

// encoding-name (string)
//
// Set the encoding name of the stream to use. This is a short-hand for
// the full caps and maps typically to the encoding-name in the RTP caps.

func (e *Ristsrc) GetEncodingName() (string, error) {
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

func (e *Ristsrc) SetEncodingName(value string) error {
	return e.Element.SetProperty("encoding-name", value)
}

// max-rtcp-bandwidth (float64)
//
// The maximum bandwidth used for RTCP as a fraction of RTP bandwdith

func (e *Ristsrc) GetMaxRtcpBandwidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("max-rtcp-bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Ristsrc) SetMaxRtcpBandwidth(value float64) error {
	return e.Element.SetProperty("max-rtcp-bandwidth", value)
}

// max-rtx-retries (uint)
//
// The maximum number of retransmission requests for a lost packet.

func (e *Ristsrc) GetMaxRtxRetries() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-rtx-retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ristsrc) SetMaxRtxRetries(value uint) error {
	return e.Element.SetProperty("max-rtx-retries", value)
}

// min-rtcp-interval (uint)
//
// The minimum interval (in ms) between two successive RTCP packets

func (e *Ristsrc) GetMinRtcpInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-rtcp-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ristsrc) SetMinRtcpInterval(value uint) error {
	return e.Element.SetProperty("min-rtcp-interval", value)
}

// multicast-iface (string)
//
// The multicast interface to use to send packets.

func (e *Ristsrc) GetMulticastIface() (string, error) {
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

func (e *Ristsrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// multicast-loopback (bool)
//
// When enabled, the packets will be received locally.

func (e *Ristsrc) GetMulticastLoopback() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("multicast-loopback")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ristsrc) SetMulticastLoopback(value bool) error {
	return e.Element.SetProperty("multicast-loopback", value)
}

// multicast-ttl (int)
//
// The multicast time-to-live parameter.

func (e *Ristsrc) GetMulticastTtl() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("multicast-ttl")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Ristsrc) SetMulticastTtl(value int) error {
	return e.Element.SetProperty("multicast-ttl", value)
}

// port (uint)
//
// The port to listen for RTP packets, the RTCP port is this value + 1. This port must be an even number.

func (e *Ristsrc) GetPort() (uint, error) {
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

func (e *Ristsrc) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

// receiver-buffer (uint)
//
// Buffering duration (in ms)

func (e *Ristsrc) GetReceiverBuffer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("receiver-buffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ristsrc) SetReceiverBuffer(value uint) error {
	return e.Element.SetProperty("receiver-buffer", value)
}

// reorder-section (uint)
//
// Time to wait before sending retransmission request (in ms)

func (e *Ristsrc) GetReorderSection() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("reorder-section")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ristsrc) SetReorderSection(value uint) error {
	return e.Element.SetProperty("reorder-section", value)
}

// stats (GstStructure)
//
// Statistic in a GstStructure named 'rist/x-receiver-stats'

func (e *Ristsrc) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// stats-update-interval (uint)
//
// The interval between 'stats' update notification (in ms) (0 disabled)

func (e *Ristsrc) GetStatsUpdateInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("stats-update-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ristsrc) SetStatsUpdateInterval(value uint) error {
	return e.Element.SetProperty("stats-update-interval", value)
}

