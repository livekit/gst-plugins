// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Send data over the network via UDP with packet destinations picked up dynamically from meta on the buffers passed
type GstDynUDPSink struct {
	*base.GstBaseSink
}

func NewDynUDPSink() (*GstDynUDPSink, error) {
	e, err := gst.NewElement("dynudpsink")
	if err != nil {
		return nil, err
	}
	return &GstDynUDPSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDynUDPSinkWithName(name string) (*GstDynUDPSink, error) {
	e, err := gst.NewElementWithName("dynudpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstDynUDPSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Close socket if passed as property on state change
// Default: true
func (e *GstDynUDPSink) SetCloseSocket(value bool) error {
	return e.Element.SetProperty("close-socket", value)
}

func (e *GstDynUDPSink) GetCloseSocket() (bool, error) {
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

// Socket to use for UDP sending. (NULL == allocate)

func (e *GstDynUDPSink) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

func (e *GstDynUDPSink) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

// Socket to use for UDPv6 sending. (NULL == allocate)

func (e *GstDynUDPSink) SetSocketV6(value interface{}) error {
	return e.Element.SetProperty("socket-v6", value)
}

func (e *GstDynUDPSink) GetSocketV6() (interface{}, error) {
	return e.Element.GetProperty("socket-v6")
}

// Address to bind the socket to
// Default: NULL
func (e *GstDynUDPSink) SetBindAddress(value string) error {
	return e.Element.SetProperty("bind-address", value)
}

func (e *GstDynUDPSink) GetBindAddress() (string, error) {
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

// Port to bind the socket to
// Default: 0, Min: 0, Max: 65535
func (e *GstDynUDPSink) SetBindPort(value int) error {
	return e.Element.SetProperty("bind-port", value)
}

func (e *GstDynUDPSink) GetBindPort() (int, error) {
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

// Send data over the network via UDP to one or multiple recipients which can be added or removed at runtime using action signals
type GstMultiUDPSink struct {
	*base.GstBaseSink
}

func NewMultiUDPSink() (*GstMultiUDPSink, error) {
	e, err := gst.NewElement("multiudpsink")
	if err != nil {
		return nil, err
	}
	return &GstMultiUDPSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewMultiUDPSinkWithName(name string) (*GstMultiUDPSink, error) {
	e, err := gst.NewElementWithName("multiudpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstMultiUDPSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// The network interface on which to join the multicast group
// Default: NULL
func (e *GstMultiUDPSink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstMultiUDPSink) GetMulticastIface() (string, error) {
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

// Quality of Service, differentiated services code point (-1 default)
// Default: -1, Min: -1, Max: 63
func (e *GstMultiUDPSink) SetQosDscp(value int) error {
	return e.Element.SetProperty("qos-dscp", value)
}

func (e *GstMultiUDPSink) GetQosDscp() (int, error) {
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

// When a destination/port pair is added multiple times, send packets multiple times as well
// Default: true
func (e *GstMultiUDPSink) SetSendDuplicates(value bool) error {
	return e.Element.SetProperty("send-duplicates", value)
}

func (e *GstMultiUDPSink) GetSendDuplicates() (bool, error) {
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

// Socket to use for UDP sending. (NULL == allocate)

func (e *GstMultiUDPSink) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

func (e *GstMultiUDPSink) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

// Address to bind the socket to
// Default: NULL
func (e *GstMultiUDPSink) SetBindAddress(value string) error {
	return e.Element.SetProperty("bind-address", value)
}

func (e *GstMultiUDPSink) GetBindAddress() (string, error) {
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

// A comma separated list of host:port pairs with destinations

func (e *GstMultiUDPSink) SetClients(value string) error {
	return e.Element.SetProperty("clients", value)
}

func (e *GstMultiUDPSink) GetClients() (string, error) {
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

// Close socket if passed as property on state change
// Default: true
func (e *GstMultiUDPSink) SetCloseSocket(value bool) error {
	return e.Element.SetProperty("close-socket", value)
}

func (e *GstMultiUDPSink) GetCloseSocket() (bool, error) {
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

// Used for setting the multicast loop parameter. TRUE = enable, FALSE = disable
// Default: true
func (e *GstMultiUDPSink) SetLoop(value bool) error {
	return e.Element.SetProperty("loop", value)
}

func (e *GstMultiUDPSink) GetLoop() (bool, error) {
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

// Automatically join/leave the multicast groups, FALSE means user has to do it himself
// Default: true
func (e *GstMultiUDPSink) SetAutoMulticast(value bool) error {
	return e.Element.SetProperty("auto-multicast", value)
}

func (e *GstMultiUDPSink) GetAutoMulticast() (bool, error) {
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

// Number of bytes received to serve to clients
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMultiUDPSink) GetBytesToServe() (uint64, error) {
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

// Forcing the use of an IPv4 socket (DEPRECATED, has no effect anymore)
// Default: false
func (e *GstMultiUDPSink) SetForceIpv4(value bool) error {
	return e.Element.SetProperty("force-ipv4", value)
}

func (e *GstMultiUDPSink) GetForceIpv4() (bool, error) {
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

// Size of the kernel send buffer in bytes, 0=default
// Default: 0, Min: 0, Max: 2147483647
func (e *GstMultiUDPSink) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstMultiUDPSink) GetBufferSize() (int, error) {
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

// Socket currently in use for UDP sending. (NULL == no socket)

func (e *GstMultiUDPSink) GetUsedSocket() (interface{}, error) {
	return e.Element.GetProperty("used-socket")
}

// Socket currently in use for UDPv6 sending. (NULL == no socket)

func (e *GstMultiUDPSink) GetUsedSocketV6() (interface{}, error) {
	return e.Element.GetProperty("used-socket-v6")
}

// Used for setting the unicast TTL parameter
// Default: 64, Min: 0, Max: 255
func (e *GstMultiUDPSink) SetTtl(value int) error {
	return e.Element.SetProperty("ttl", value)
}

func (e *GstMultiUDPSink) GetTtl() (int, error) {
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

// Used for setting the multicast TTL parameter
// Default: 1, Min: 0, Max: 255
func (e *GstMultiUDPSink) SetTtlMc(value int) error {
	return e.Element.SetProperty("ttl-mc", value)
}

func (e *GstMultiUDPSink) GetTtlMc() (int, error) {
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

// Port to bind the socket to
// Default: 0, Min: 0, Max: 65535
func (e *GstMultiUDPSink) SetBindPort(value int) error {
	return e.Element.SetProperty("bind-port", value)
}

func (e *GstMultiUDPSink) GetBindPort() (int, error) {
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

// Total number of bytes sent to all clients
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMultiUDPSink) GetBytesServed() (uint64, error) {
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

// Socket to use for UDPv6 sending. (NULL == allocate)

func (e *GstMultiUDPSink) SetSocketV6(value interface{}) error {
	return e.Element.SetProperty("socket-v6", value)
}

func (e *GstMultiUDPSink) GetSocketV6() (interface{}, error) {
	return e.Element.GetProperty("socket-v6")
}

// Send data over the network via UDP
type GstUDPSink struct {
	*GstMultiUDPSink
}

func NewUDPSink() (*GstUDPSink, error) {
	e, err := gst.NewElement("udpsink")
	if err != nil {
		return nil, err
	}
	return &GstUDPSink{GstMultiUDPSink: &GstMultiUDPSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewUDPSinkWithName(name string) (*GstUDPSink, error) {
	e, err := gst.NewElementWithName("udpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstUDPSink{GstMultiUDPSink: &GstMultiUDPSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// The port to send the packets to
// Default: 5004, Min: 0, Max: 65535
func (e *GstUDPSink) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstUDPSink) GetPort() (int, error) {
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

// The host/IP/Multicast group to send the packets to
// Default: localhost
func (e *GstUDPSink) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

func (e *GstUDPSink) GetHost() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("host")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Receive data over the network via UDP
type GstUDPSrc struct {
	*base.GstPushSrc
}

func NewUDPSrc() (*GstUDPSrc, error) {
	e, err := gst.NewElement("udpsrc")
	if err != nil {
		return nil, err
	}
	return &GstUDPSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewUDPSrcWithName(name string) (*GstUDPSrc, error) {
	e, err := gst.NewElementWithName("udpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstUDPSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Maximum expected packet size. This directly defines the allocationsize of the receive buffer pool.
// Default: 1492, Min: 0, Max: 2147483647
func (e *GstUDPSrc) SetMtu(value uint) error {
	return e.Element.SetProperty("mtu", value)
}

func (e *GstUDPSrc) GetMtu() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("mtu")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The Address of multicast group to join. (DEPRECATED: Use address property instead)
// Default: 0.0.0.0
func (e *GstUDPSrc) SetMulticastGroup(value string) error {
	return e.Element.SetProperty("multicast-group", value)
}

func (e *GstUDPSrc) GetMulticastGroup() (string, error) {
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

// The network interface on which to join the multicast group.This allows multiple interfaces separated by comma. ("eth0,eth1")
// Default: NULL
func (e *GstUDPSrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstUDPSrc) GetMulticastIface() (string, error) {
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

// Whether to retrieve the sender address and add it to buffers as meta. Disabling this might result in minor performance improvements in certain scenarios
// Default: true
func (e *GstUDPSrc) SetRetrieveSenderAddress(value bool) error {
	return e.Element.SetProperty("retrieve-sender-address", value)
}

func (e *GstUDPSrc) GetRetrieveSenderAddress() (bool, error) {
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

// URI in the form of udp://multicast_group:port
// Default: udp://0.0.0.0:5004
func (e *GstUDPSrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstUDPSrc) GetUri() (string, error) {
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

// Socket currently in use for UDP reception. (NULL = no socket)

func (e *GstUDPSrc) GetUsedSocket() (interface{}, error) {
	return e.Element.GetProperty("used-socket")
}

// Used for setting the multicast loop parameter. TRUE = enable, FALSE = disable
// Default: true
func (e *GstUDPSrc) SetLoop(value bool) error {
	return e.Element.SetProperty("loop", value)
}

func (e *GstUDPSrc) GetLoop() (bool, error) {
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

// Address to receive packets for. This is equivalent to the multicast-group property for now
// Default: 0.0.0.0
func (e *GstUDPSrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstUDPSrc) GetAddress() (string, error) {
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

// Close socket if passed as property on state change
// Default: true
func (e *GstUDPSrc) SetCloseSocket(value bool) error {
	return e.Element.SetProperty("close-socket", value)
}

func (e *GstUDPSrc) GetCloseSocket() (bool, error) {
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

// Used for adding alternative timestamp using SO_TIMESTAMP.
// Default: disabled (0)
func (e *GstUDPSrc) SetSocketTimestamp(value GstSocketTimestampMode) error {
	e.Element.SetArg("socket-timestamp", string(value))
	return nil
}

func (e *GstUDPSrc) GetSocketTimestamp() (GstSocketTimestampMode, error) {
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

// Post a message after timeout nanoseconds (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstUDPSrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstUDPSrc) GetTimeout() (uint64, error) {
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

// Size of the kernel receive buffer in bytes, 0=default
// Default: 0, Min: 0, Max: 2147483647
func (e *GstUDPSrc) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstUDPSrc) GetBufferSize() (int, error) {
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

// The caps of the source pad

func (e *GstUDPSrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstUDPSrc) GetCaps() (*gst.Caps, error) {
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

// List of source to receive the stream with '+' (positive filter) or '-' (negative filter, ignored for now) prefix (e.g., "+SOURCE0+SOURCE1+SOURCE2"). Alternatively, user can use URI query with the key value "multicast-source"
// Default: NULL
func (e *GstUDPSrc) SetMulticastSource(value string) error {
	return e.Element.SetProperty("multicast-source", value)
}

func (e *GstUDPSrc) GetMulticastSource() (string, error) {
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

// The port to receive the packets from, 0=allocate
// Default: 5004, Min: 0, Max: 65535
func (e *GstUDPSrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstUDPSrc) GetPort() (int, error) {
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

// Enable reuse of the port
// Default: true
func (e *GstUDPSrc) SetReuse(value bool) error {
	return e.Element.SetProperty("reuse", value)
}

func (e *GstUDPSrc) GetReuse() (bool, error) {
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

// number of bytes to skip for each udp packet
// Default: 0, Min: 0, Max: 2147483647
func (e *GstUDPSrc) SetSkipFirstBytes(value int) error {
	return e.Element.SetProperty("skip-first-bytes", value)
}

func (e *GstUDPSrc) GetSkipFirstBytes() (int, error) {
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

// Socket to use for UDP reception. (NULL == allocate)

func (e *GstUDPSrc) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

func (e *GstUDPSrc) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

// Automatically join/leave multicast groups
// Default: true
func (e *GstUDPSrc) SetAutoMulticast(value bool) error {
	return e.Element.SetProperty("auto-multicast", value)
}

func (e *GstUDPSrc) GetAutoMulticast() (bool, error) {
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

type GstSocketTimestampMode string

const (
	GstSocketTimestampModeDisabled GstSocketTimestampMode = "disabled" // Disable additional timestamps
	GstSocketTimestampModeRealtime GstSocketTimestampMode = "realtime" // Timestamp with realtime clock (nsec resolution, may not be monotonic)
)
