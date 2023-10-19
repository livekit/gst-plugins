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

type Udpsrc struct {
	*base.GstPushSrc
}

func NewUdpsrc() (*Udpsrc, error) {
	e, err := gst.NewElement("udpsrc")
	if err != nil {
		return nil, err
	}
	return &Udpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewUdpsrcWithName(name string) (*Udpsrc, error) {
	e, err := gst.NewElementWithName("udpsrc", name)
	if err != nil {
		return nil, err
	}
	return &Udpsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// address (string)
//
// Address to receive packets for. This is equivalent to the multicast-group property for now

func (e *Udpsrc) GetAddress() (string, error) {
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

func (e *Udpsrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

// auto-multicast (bool)
//
// Automatically join/leave multicast groups

func (e *Udpsrc) GetAutoMulticast() (bool, error) {
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

func (e *Udpsrc) SetAutoMulticast(value bool) error {
	return e.Element.SetProperty("auto-multicast", value)
}

// buffer-size (int)
//
// Size of the kernel receive buffer in bytes, 0=default

func (e *Udpsrc) GetBufferSize() (int, error) {
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

func (e *Udpsrc) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// caps (GstCaps)
//
// The caps of the source pad

func (e *Udpsrc) GetCaps() (*gst.Caps, error) {
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

func (e *Udpsrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// close-socket (bool)
//
// Close socket if passed as property on state change

func (e *Udpsrc) GetCloseSocket() (bool, error) {
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

func (e *Udpsrc) SetCloseSocket(value bool) error {
	return e.Element.SetProperty("close-socket", value)
}

// loop (bool)
//
// Can be used to disable multicast loopback.

func (e *Udpsrc) GetLoop() (bool, error) {
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

func (e *Udpsrc) SetLoop(value bool) error {
	return e.Element.SetProperty("loop", value)
}

// mtu (uint)
//
// Maximum expected packet size. This directly defines the allocation
// size of the receive buffer pool.

// multicast-group (string)
//
// The Address of multicast group to join. (DEPRECATED: Use address property instead)

func (e *Udpsrc) GetMulticastGroup() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("multicast-group")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Udpsrc) SetMulticastGroup(value string) error {
	return e.Element.SetProperty("multicast-group", value)
}

// multicast-iface (string)
//
// The network interface on which to join the multicast group.This allows multiple interfaces separated by comma. ("eth0,eth1")

func (e *Udpsrc) GetMulticastIface() (string, error) {
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

func (e *Udpsrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// multicast-source (string)
//
// List of source to receive the stream from. (IGMPv3 SSM RFC 4604)

func (e *Udpsrc) GetMulticastSource() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("multicast-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Udpsrc) SetMulticastSource(value string) error {
	return e.Element.SetProperty("multicast-source", value)
}

// port (int)
//
// The port to receive the packets from, 0=allocate

func (e *Udpsrc) GetPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Udpsrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

// retrieve-sender-address (bool)
//
// Whether to retrieve the sender address and add it to the buffers as
// meta. Disabling this might result in minor performance improvements
// in certain scenarios.

func (e *Udpsrc) GetRetrieveSenderAddress() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("retrieve-sender-address")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Udpsrc) SetRetrieveSenderAddress(value bool) error {
	return e.Element.SetProperty("retrieve-sender-address", value)
}

// reuse (bool)
//
// Enable reuse of the port

func (e *Udpsrc) GetReuse() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reuse")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Udpsrc) SetReuse(value bool) error {
	return e.Element.SetProperty("reuse", value)
}

// skip-first-bytes (int)
//
// number of bytes to skip for each udp packet

func (e *Udpsrc) GetSkipFirstBytes() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("skip-first-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Udpsrc) SetSkipFirstBytes(value int) error {
	return e.Element.SetProperty("skip-first-bytes", value)
}

// socket (GstGsocket)
//
// Socket to use for UDP reception. (NULL == allocate)

func (e *Udpsrc) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

func (e *Udpsrc) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

// socket-timestamp (GstSocketTimestampMode)
//
// Can be used to read the timestamp on incoming buffers using socket
// control messages and set as the DTS.

func (e *Udpsrc) GetSocketTimestamp() (GstSocketTimestampMode, error) {
	var value GstSocketTimestampMode
	var ok bool
	v, err := e.Element.GetProperty("socket-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSocketTimestampMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSocketTimestampMode")
	}
	return value, nil
}

func (e *Udpsrc) SetSocketTimestamp(value GstSocketTimestampMode) error {
	e.Element.SetArg("socket-timestamp", string(value))
	return nil
}

// timeout (uint64)
//
// Post a message after timeout nanoseconds (0 = disabled)

func (e *Udpsrc) GetTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Udpsrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

// uri (string)
//
// URI in the form of udp://multicast_group:port

func (e *Udpsrc) GetUri() (string, error) {
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

func (e *Udpsrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// used-socket (GstGsocket)
//
// Socket currently in use for UDP reception. (NULL = no socket)

func (e *Udpsrc) GetUsedSocket() (interface{}, error) {
	return e.Element.GetProperty("used-socket")
}

// ----- Constants -----

type GstSocketTimestampMode string

const (
	GstSocketTimestampModeDisabled GstSocketTimestampMode = "disabled" // disabled (0) – Disable additional timestamps
	GstSocketTimestampModeRealtime GstSocketTimestampMode = "realtime" // realtime (1) – Timestamp with realtime clock (nsec resolution, may not be monotonic)
)

