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

	"github.com/livekit/gstplugins/base"
)

type Multiudpsink struct {
	*base.GstBaseSink
}

func NewMultiudpsink() (*Multiudpsink, error) {
	e, err := gst.NewElement("multiudpsink")
	if err != nil {
		return nil, err
	}
	return &Multiudpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewMultiudpsinkWithName(name string) (*Multiudpsink, error) {
	e, err := gst.NewElementWithName("multiudpsink", name)
	if err != nil {
		return nil, err
	}
	return &Multiudpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// auto-multicast (bool)
//
// Automatically join/leave the multicast groups, FALSE means user has to do it himself

func (e *Multiudpsink) GetAutoMulticast() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-multicast")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiudpsink) SetAutoMulticast(value bool) error {
	return e.Element.SetProperty("auto-multicast", value)
}

// bind-address (string)
//
// Address to bind the socket to

func (e *Multiudpsink) GetBindAddress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("bind-address")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Multiudpsink) SetBindAddress(value string) error {
	return e.Element.SetProperty("bind-address", value)
}

// bind-port (int)
//
// Port to bind the socket to

func (e *Multiudpsink) GetBindPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bind-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multiudpsink) SetBindPort(value int) error {
	return e.Element.SetProperty("bind-port", value)
}

// buffer-size (int)
//
// Size of the kernel send buffer in bytes, 0=default

func (e *Multiudpsink) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multiudpsink) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// bytes-served (uint64)
//
// Total number of bytes sent to all clients

func (e *Multiudpsink) GetBytesServed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("bytes-served")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// bytes-to-serve (uint64)
//
// Number of bytes received to serve to clients

func (e *Multiudpsink) GetBytesToServe() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("bytes-to-serve")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// clients (string)
//
// A comma separated list of host:port pairs with destinations

func (e *Multiudpsink) GetClients() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("clients")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Multiudpsink) SetClients(value string) error {
	return e.Element.SetProperty("clients", value)
}

// close-socket (bool)
//
// Close socket if passed as property on state change

func (e *Multiudpsink) GetCloseSocket() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("close-socket")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiudpsink) SetCloseSocket(value bool) error {
	return e.Element.SetProperty("close-socket", value)
}

// force-ipv4 (bool)
//
// Forcing the use of an IPv4 socket (DEPRECATED, has no effect anymore)

func (e *Multiudpsink) GetForceIpv4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-ipv4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiudpsink) SetForceIpv4(value bool) error {
	return e.Element.SetProperty("force-ipv4", value)
}

// loop (bool)
//
// Used for setting the multicast loop parameter. TRUE = enable, FALSE = disable

func (e *Multiudpsink) GetLoop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiudpsink) SetLoop(value bool) error {
	return e.Element.SetProperty("loop", value)
}

// multicast-iface (string)
//
// The network interface on which to join the multicast group

func (e *Multiudpsink) GetMulticastIface() (string, error) {
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

func (e *Multiudpsink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// qos-dscp (int)
//
// Quality of Service, differentiated services code point (-1 default)

func (e *Multiudpsink) GetQosDscp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qos-dscp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multiudpsink) SetQosDscp(value int) error {
	return e.Element.SetProperty("qos-dscp", value)
}

// send-duplicates (bool)
//
// When a destination/port pair is added multiple times, send packets multiple times as well

func (e *Multiudpsink) GetSendDuplicates() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-duplicates")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiudpsink) SetSendDuplicates(value bool) error {
	return e.Element.SetProperty("send-duplicates", value)
}

// socket (GstGsocket)
//
// Socket to use for UDP sending. (NULL == allocate)

func (e *Multiudpsink) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

func (e *Multiudpsink) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

// socket-v6 (GstGsocket)
//
// Socket to use for UDPv6 sending. (NULL == allocate)

func (e *Multiudpsink) GetSocketV6() (interface{}, error) {
	return e.Element.GetProperty("socket-v6")
}

func (e *Multiudpsink) SetSocketV6(value interface{}) error {
	return e.Element.SetProperty("socket-v6", value)
}

// ttl (int)
//
// Used for setting the unicast TTL parameter

func (e *Multiudpsink) GetTtl() (int, error) {
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

func (e *Multiudpsink) SetTtl(value int) error {
	return e.Element.SetProperty("ttl", value)
}

// ttl-mc (int)
//
// Used for setting the multicast TTL parameter

func (e *Multiudpsink) GetTtlMc() (int, error) {
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

func (e *Multiudpsink) SetTtlMc(value int) error {
	return e.Element.SetProperty("ttl-mc", value)
}

// used-socket (GstGsocket)
//
// Socket currently in use for UDP sending. (NULL == no socket)

func (e *Multiudpsink) GetUsedSocket() (interface{}, error) {
	return e.Element.GetProperty("used-socket")
}

// used-socket-v6 (GstGsocket)
//
// Socket currently in use for UDPv6 sending. (NULL == no socket)

func (e *Multiudpsink) GetUsedSocketV6() (interface{}, error) {
	return e.Element.GetProperty("used-socket-v6")
}

