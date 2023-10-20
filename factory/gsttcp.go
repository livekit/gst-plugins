// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Receive data from a socket
type GstSocketSrc struct {
	*base.GstPushSrc
}

func NewSocketSrc() (*GstSocketSrc, error) {
	e, err := gst.NewElement("socketsrc")
	if err != nil {
		return nil, err
	}
	return &GstSocketSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewSocketSrcWithName(name string) (*GstSocketSrc, error) {
	e, err := gst.NewElementWithName("socketsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSocketSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// The caps of the source pad

func (e *GstSocketSrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstSocketSrc) GetCaps() (*gst.Caps, error) {
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

// If GstNetworkMessage events should be handled
// Default: false
func (e *GstSocketSrc) SetSendMessages(value bool) error {
	return e.Element.SetProperty("send-messages", value)
}

func (e *GstSocketSrc) GetSendMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The socket to receive packets from

func (e *GstSocketSrc) SetSocket(value interface{}) error {
	return e.Element.SetProperty("socket", value)
}

func (e *GstSocketSrc) GetSocket() (interface{}, error) {
	return e.Element.GetProperty("socket")
}

// Send data as a client over the network via TCP
type GstTCPClientSink struct {
	*base.GstBaseSink
}

func NewTCPClientSink() (*GstTCPClientSink, error) {
	e, err := gst.NewElement("tcpclientsink")
	if err != nil {
		return nil, err
	}
	return &GstTCPClientSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewTCPClientSinkWithName(name string) (*GstTCPClientSink, error) {
	e, err := gst.NewElementWithName("tcpclientsink", name)
	if err != nil {
		return nil, err
	}
	return &GstTCPClientSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// The host/IP to send the packets to
// Default: localhost
func (e *GstTCPClientSink) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

func (e *GstTCPClientSink) GetHost() (string, error) {
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

// The port to send the packets to
// Default: 4953, Min: 0, Max: 65535
func (e *GstTCPClientSink) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstTCPClientSink) GetPort() (int, error) {
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

// Receive data as a client over the network via TCP
type GstTCPClientSrc struct {
	*base.GstPushSrc
}

func NewTCPClientSrc() (*GstTCPClientSrc, error) {
	e, err := gst.NewElement("tcpclientsrc")
	if err != nil {
		return nil, err
	}
	return &GstTCPClientSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewTCPClientSrcWithName(name string) (*GstTCPClientSrc, error) {
	e, err := gst.NewElementWithName("tcpclientsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstTCPClientSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Retrieve a statistics structure
// Default: GstTCPClientSrcStats, bytes-received=(guint64)0;
func (e *GstTCPClientSrc) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Value in seconds to timeout a blocking I/O. 0 = No timeout.
// Default: 0, Min: 0, Max: -1
func (e *GstTCPClientSrc) SetTimeout(value uint) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstTCPClientSrc) GetTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The host IP address to receive packets from
// Default: localhost
func (e *GstTCPClientSrc) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

func (e *GstTCPClientSrc) GetHost() (string, error) {
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

// The port to receive packets from
// Default: 4953, Min: 0, Max: 65535
func (e *GstTCPClientSrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstTCPClientSrc) GetPort() (int, error) {
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

// Send data as a server over the network via TCP
type GstTCPServerSink struct {
	*GstMultiSocketSink
}

func NewTCPServerSink() (*GstTCPServerSink, error) {
	e, err := gst.NewElement("tcpserversink")
	if err != nil {
		return nil, err
	}
	return &GstTCPServerSink{GstMultiSocketSink: &GstMultiSocketSink{GstMultiHandleSink: &GstMultiHandleSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewTCPServerSinkWithName(name string) (*GstTCPServerSink, error) {
	e, err := gst.NewElementWithName("tcpserversink", name)
	if err != nil {
		return nil, err
	}
	return &GstTCPServerSink{GstMultiSocketSink: &GstMultiSocketSink{GstMultiHandleSink: &GstMultiHandleSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// The port number the socket is currently bound to
// Default: 0, Min: 0, Max: 65535
func (e *GstTCPServerSink) GetCurrentPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("current-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The host/IP to listen on
// Default: localhost
func (e *GstTCPServerSink) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

func (e *GstTCPServerSink) GetHost() (string, error) {
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

// The port to listen to (0=random available port)
// Default: 4953, Min: 0, Max: 65535
func (e *GstTCPServerSink) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstTCPServerSink) GetPort() (int, error) {
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

// Receive data as a server over the network via TCP
type GstTCPServerSrc struct {
	*base.GstPushSrc
}

func NewTCPServerSrc() (*GstTCPServerSrc, error) {
	e, err := gst.NewElement("tcpserversrc")
	if err != nil {
		return nil, err
	}
	return &GstTCPServerSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewTCPServerSrcWithName(name string) (*GstTCPServerSrc, error) {
	e, err := gst.NewElementWithName("tcpserversrc", name)
	if err != nil {
		return nil, err
	}
	return &GstTCPServerSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Retrieve a statistics structure
// Default: GstTCPServerSrcStats, bytes-received=(guint64)0;
func (e *GstTCPServerSrc) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// The port number the socket is currently bound to
// Default: 0, Min: 0, Max: 65535
func (e *GstTCPServerSrc) GetCurrentPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("current-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The hostname to listen as
// Default: localhost
func (e *GstTCPServerSrc) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

func (e *GstTCPServerSrc) GetHost() (string, error) {
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

// The port to listen to (0=random available port)
// Default: 4953, Min: 0, Max: 65535
func (e *GstTCPServerSrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstTCPServerSrc) GetPort() (int, error) {
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

// Send data to multiple filedescriptors
type GstMultiFdSink struct {
	*GstMultiHandleSink
}

func NewMultiFdSink() (*GstMultiFdSink, error) {
	e, err := gst.NewElement("multifdsink")
	if err != nil {
		return nil, err
	}
	return &GstMultiFdSink{GstMultiHandleSink: &GstMultiHandleSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewMultiFdSinkWithName(name string) (*GstMultiFdSink, error) {
	e, err := gst.NewElementWithName("multifdsink", name)
	if err != nil {
		return nil, err
	}
	return &GstMultiFdSink{GstMultiHandleSink: &GstMultiHandleSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// Handle client reads and discard the data
// Default: true
func (e *GstMultiFdSink) SetHandleRead(value bool) error {
	return e.Element.SetProperty("handle-read", value)
}

func (e *GstMultiFdSink) GetHandleRead() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("handle-read")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send data to multiple sockets
type GstMultiSocketSink struct {
	*GstMultiHandleSink
}

func NewMultiSocketSink() (*GstMultiSocketSink, error) {
	e, err := gst.NewElement("multisocketsink")
	if err != nil {
		return nil, err
	}
	return &GstMultiSocketSink{GstMultiHandleSink: &GstMultiHandleSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

func NewMultiSocketSinkWithName(name string) (*GstMultiSocketSink, error) {
	e, err := gst.NewElementWithName("multisocketsink", name)
	if err != nil {
		return nil, err
	}
	return &GstMultiSocketSink{GstMultiHandleSink: &GstMultiHandleSink{GstBaseSink: &base.GstBaseSink{Element: e}}}, nil
}

// If GstNetworkMessageDispatched events should be pushed
// Default: false
func (e *GstMultiSocketSink) SetSendDispatched(value bool) error {
	return e.Element.SetProperty("send-dispatched", value)
}

func (e *GstMultiSocketSink) GetSendDispatched() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-dispatched")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If GstNetworkMessage events should be pushed
// Default: false
func (e *GstMultiSocketSink) SetSendMessages(value bool) error {
	return e.Element.SetProperty("send-messages", value)
}

func (e *GstMultiSocketSink) GetSendMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstMultiHandleSink struct {
	*base.GstBaseSink
}

// max number of buffers to queue for a client (-1 = no limit)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMultiHandleSink) SetBuffersMax(value int) error {
	return e.Element.SetProperty("buffers-max", value)
}

func (e *GstMultiHandleSink) GetBuffersMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffers-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The format of the burst units (when sync-method is burst[[-with]-keyframe])
// Default: undefined (0)
func (e *GstMultiHandleSink) SetBurstFormat(value *gst.Format) error {
	return e.Element.SetProperty("burst-format", value)
}

func (e *GstMultiHandleSink) GetBurstFormat() (*gst.Format, error) {
	var value *gst.Format
	var ok bool
	v, err := e.Element.GetProperty("burst-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Format)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Format")
	}
	return value, nil
}

// The amount of burst expressed in burst-format
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMultiHandleSink) SetBurstValue(value uint64) error {
	return e.Element.SetProperty("burst-value", value)
}

func (e *GstMultiHandleSink) GetBurstValue() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("burst-value")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Total number of bytes send to all clients
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMultiHandleSink) GetBytesServed() (uint64, error) {
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

// max number of units to queue (-1 = no limit)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstMultiHandleSink) SetUnitsMax(value int64) error {
	return e.Element.SetProperty("units-max", value)
}

func (e *GstMultiHandleSink) GetUnitsMax() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("units-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Recover client when going over this limit (-1 = no limit)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstMultiHandleSink) SetUnitsSoftMax(value int64) error {
	return e.Element.SetProperty("units-soft-max", value)
}

func (e *GstMultiHandleSink) GetUnitsSoftMax() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("units-soft-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Quality of Service, differentiated services code point (-1 default)
// Default: -1, Min: -1, Max: 63
func (e *GstMultiHandleSink) SetQosDscp(value int) error {
	return e.Element.SetProperty("qos-dscp", value)
}

func (e *GstMultiHandleSink) GetQosDscp() (int, error) {
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

// Maximum inactivity timeout in nanoseconds for a client (0 = no limit)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMultiHandleSink) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstMultiHandleSink) GetTimeout() (uint64, error) {
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

// The unit to measure the max/soft-max/queued properties
// Default: buffers (4)
func (e *GstMultiHandleSink) SetUnitFormat(value *gst.Format) error {
	return e.Element.SetProperty("unit-format", value)
}

func (e *GstMultiHandleSink) GetUnitFormat() (*gst.Format, error) {
	var value *gst.Format
	var ok bool
	v, err := e.Element.GetProperty("unit-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Format)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Format")
	}
	return value, nil
}

// min number of buffers to queue (-1 = as few as possible)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMultiHandleSink) SetBuffersMin(value int) error {
	return e.Element.SetProperty("buffers-min", value)
}

func (e *GstMultiHandleSink) GetBuffersMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffers-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of buffers currently queued
// Default: 0, Min: 0, Max: -1
func (e *GstMultiHandleSink) GetBuffersQueued() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffers-queued")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// min number of bytes to queue (-1 = as little as possible)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMultiHandleSink) SetBytesMin(value int) error {
	return e.Element.SetProperty("bytes-min", value)
}

func (e *GstMultiHandleSink) GetBytesMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bytes-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The current number of client handles
// Default: 0, Min: 0, Max: -1
func (e *GstMultiHandleSink) GetNumHandles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-handles")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// How to recover when client reaches the soft max
// Default: none (0)
func (e *GstMultiHandleSink) SetRecoverPolicy(value GstMultiHandleSinkRecoverPolicy) error {
	e.Element.SetArg("recover-policy", string(value))
	return nil
}

func (e *GstMultiHandleSink) GetRecoverPolicy() (GstMultiHandleSinkRecoverPolicy, error) {
	var value GstMultiHandleSinkRecoverPolicy
	var ok bool
	v, err := e.Element.GetProperty("recover-policy")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMultiHandleSinkRecoverPolicy)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMultiHandleSinkRecoverPolicy")
	}
	return value, nil
}

// min amount of time to queue (in nanoseconds) (-1 = as little as possible)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstMultiHandleSink) SetTimeMin(value int64) error {
	return e.Element.SetProperty("time-min", value)
}

func (e *GstMultiHandleSink) GetTimeMin() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("time-min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Recover client when going over this limit (-1 = no limit)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstMultiHandleSink) SetBuffersSoftMax(value int) error {
	return e.Element.SetProperty("buffers-soft-max", value)
}

func (e *GstMultiHandleSink) GetBuffersSoftMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffers-soft-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of bytes received to serve to clients
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstMultiHandleSink) GetBytesToServe() (uint64, error) {
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

// Resend the streamheader if it changes in the caps
// Default: true
func (e *GstMultiHandleSink) SetResendStreamheader(value bool) error {
	return e.Element.SetProperty("resend-streamheader", value)
}

func (e *GstMultiHandleSink) GetResendStreamheader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("resend-streamheader")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// How to sync new clients to the stream
// Default: latest (0)
func (e *GstMultiHandleSink) SetSyncMethod(value GstMultiHandleSinkSyncMethod) error {
	e.Element.SetArg("sync-method", string(value))
	return nil
}

func (e *GstMultiHandleSink) GetSyncMethod() (GstMultiHandleSinkSyncMethod, error) {
	var value GstMultiHandleSinkSyncMethod
	var ok bool
	v, err := e.Element.GetProperty("sync-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMultiHandleSinkSyncMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMultiHandleSinkSyncMethod")
	}
	return value, nil
}

type GstMultiHandleSinkClientStatus string

const (
	GstMultiHandleSinkClientStatusOk        GstMultiHandleSinkClientStatus = "ok"        // ok
	GstMultiHandleSinkClientStatusClosed    GstMultiHandleSinkClientStatus = "closed"    // Closed
	GstMultiHandleSinkClientStatusRemoved   GstMultiHandleSinkClientStatus = "removed"   // Removed
	GstMultiHandleSinkClientStatusSlow      GstMultiHandleSinkClientStatus = "slow"      // Too slow
	GstMultiHandleSinkClientStatusError     GstMultiHandleSinkClientStatus = "error"     // Error
	GstMultiHandleSinkClientStatusDuplicate GstMultiHandleSinkClientStatus = "duplicate" // Duplicate
	GstMultiHandleSinkClientStatusFlushing  GstMultiHandleSinkClientStatus = "flushing"  // Flushing
)

type GstMultiHandleSinkRecoverPolicy string

const (
	GstMultiHandleSinkRecoverPolicyNone      GstMultiHandleSinkRecoverPolicy = "none"       // Do not try to recover
	GstMultiHandleSinkRecoverPolicyLatest    GstMultiHandleSinkRecoverPolicy = "latest"     // Resync client to latest buffer
	GstMultiHandleSinkRecoverPolicySoftLimit GstMultiHandleSinkRecoverPolicy = "soft-limit" // Resync client to soft limit
	GstMultiHandleSinkRecoverPolicyKeyframe  GstMultiHandleSinkRecoverPolicy = "keyframe"   // Resync client to most recent keyframe
)

type GstMultiHandleSinkSyncMethod string

const (
	GstMultiHandleSinkSyncMethodLatest            GstMultiHandleSinkSyncMethod = "latest"              // Serve starting from the latest buffer
	GstMultiHandleSinkSyncMethodNextKeyframe      GstMultiHandleSinkSyncMethod = "next-keyframe"       // Serve starting from the next keyframe
	GstMultiHandleSinkSyncMethodLatestKeyframe    GstMultiHandleSinkSyncMethod = "latest-keyframe"     // Serve everything since the latest keyframe (burst)
	GstMultiHandleSinkSyncMethodBurst             GstMultiHandleSinkSyncMethod = "burst"               // Serve burst-value data to client
	GstMultiHandleSinkSyncMethodBurstKeyframe     GstMultiHandleSinkSyncMethod = "burst-keyframe"      // Serve burst-value data starting on a keyframe
	GstMultiHandleSinkSyncMethodBurstWithKeyframe GstMultiHandleSinkSyncMethod = "burst-with-keyframe" // Serve burst-value data preferably starting on a keyframe
)
