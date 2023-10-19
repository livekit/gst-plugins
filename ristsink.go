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

type Ristsink struct {
	Element *gst.Element
}

func NewRistsink() (*Ristsink, error) {
	e, err := gst.NewElement("ristsink")
	if err != nil {
		return nil, err
	}
	return &Ristsink{Element: e}, nil
}

func NewRistsinkWithName(name string) (*Ristsink, error) {
	e, err := gst.NewElementWithName("ristsink", name)
	if err != nil {
		return nil, err
	}
	return &Ristsink{Element: e}, nil
}

// ----- Properties -----

// address (string)
//
// Address to send packets to (can be IPv4 or IPv6).

func (e *Ristsink) GetAddress() (string, error) {
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

func (e *Ristsink) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

// bonding-addresses (string)
//
// Comma (,) separated list of 

func (e *Ristsink) GetBondingAddresses() (string, error) {
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

func (e *Ristsink) SetBondingAddresses(value string) error {
	return e.Element.SetProperty("bonding-addresses", value)
}

// bonding-method (GstRistBondingMethodType)
//
// Defines the bonding method to use.

func (e *Ristsink) GetBondingMethod() (GstRistBondingMethodType, error) {
	var value GstRistBondingMethodType
	var ok bool
	v, err := e.Element.GetProperty("bonding-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRistBondingMethodType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRistBondingMethodType")
	}
	return value, nil
}

func (e *Ristsink) SetBondingMethod(value GstRistBondingMethodType) error {
	e.Element.SetArg("bonding-method", string(value))
	return nil
}

// cname (string)
//
// Set the CNAME in the SDES block of the sender report.

func (e *Ristsink) GetCname() (string, error) {
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

func (e *Ristsink) SetCname(value string) error {
	return e.Element.SetProperty("cname", value)
}

// dispatcher (GstElement)
//
// An element that takes care of multi-plexing bonded links. When set "bonding-method" is ignored.

func (e *Ristsink) GetDispatcher() (interface{}, error) {
	return e.Element.GetProperty("dispatcher")
}

func (e *Ristsink) SetDispatcher(value interface{}) error {
	return e.Element.SetProperty("dispatcher", value)
}

// drop-null-ts-packets (bool)
//
// Drop null MPEG-TS packet and replace them with a custom header extension.

func (e *Ristsink) GetDropNullTsPackets() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-null-ts-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ristsink) SetDropNullTsPackets(value bool) error {
	return e.Element.SetProperty("drop-null-ts-packets", value)
}

// max-rtcp-bandwidth (float64)
//
// The maximum bandwidth used for RTCP as a fraction of RTP bandwdith

func (e *Ristsink) GetMaxRtcpBandwidth() (float64, error) {
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

func (e *Ristsink) SetMaxRtcpBandwidth(value float64) error {
	return e.Element.SetProperty("max-rtcp-bandwidth", value)
}

// min-rtcp-interval (uint)
//
// The minimum interval (in ms) between two regular successive RTCP packets.

func (e *Ristsink) GetMinRtcpInterval() (uint, error) {
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

func (e *Ristsink) SetMinRtcpInterval(value uint) error {
	return e.Element.SetProperty("min-rtcp-interval", value)
}

// multicast-iface (string)
//
// The multicast interface to use to send packets.

func (e *Ristsink) GetMulticastIface() (string, error) {
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

func (e *Ristsink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// multicast-loopback (bool)
//
// When enabled, the packet will be received locally.

func (e *Ristsink) GetMulticastLoopback() (bool, error) {
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

func (e *Ristsink) SetMulticastLoopback(value bool) error {
	return e.Element.SetProperty("multicast-loopback", value)
}

// multicast-ttl (int)
//
// The multicast time-to-live parameter.

func (e *Ristsink) GetMulticastTtl() (int, error) {
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

func (e *Ristsink) SetMulticastTtl(value int) error {
	return e.Element.SetProperty("multicast-ttl", value)
}

// port (uint)
//
// The port RTP packets will be sent, the RTCP port is this value + 1. This port must be an even number.

func (e *Ristsink) GetPort() (uint, error) {
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

func (e *Ristsink) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

// sender-buffer (uint)
//
// Size of the retransmission queue (in ms)

func (e *Ristsink) GetSenderBuffer() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sender-buffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ristsink) SetSenderBuffer(value uint) error {
	return e.Element.SetProperty("sender-buffer", value)
}

// sequence-number-extension (bool)
//
// Add sequence number extension to packets.

func (e *Ristsink) GetSequenceNumberExtension() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sequence-number-extension")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ristsink) SetSequenceNumberExtension(value bool) error {
	return e.Element.SetProperty("sequence-number-extension", value)
}

// stats (GstStructure)
//
// Statistic in a GstStructure named 'rist/x-sender-stats'

func (e *Ristsink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// stats-update-interval (uint)
//
// The interval between 'stats' update notification (in ms) (0 disabled)

func (e *Ristsink) GetStatsUpdateInterval() (uint, error) {
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

func (e *Ristsink) SetStatsUpdateInterval(value uint) error {
	return e.Element.SetProperty("stats-update-interval", value)
}

// ----- Constants -----

type GstRistBondingMethodType string

const (
	GstRistBondingMethodTypeBroadcast GstRistBondingMethodType = "broadcast" // broadcast (0) – GST_RIST_BONDING_METHOD_BROADCAST
	GstRistBondingMethodTypeRoundRobin GstRistBondingMethodType = "round-robin" // round-robin (1) – GST_RIST_BONDING_METHOD_ROUND_ROBIN
)

