// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Simple RTP sink
type GstRtpSink struct {
	*gst.Bin
}

func NewRtpSink() (*GstRtpSink, error) {
	e, err := gst.NewElement("rtpsink")
	if err != nil {
		return nil, err
	}
	return &GstRtpSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewRtpSinkWithName(name string) (*GstRtpSink, error) {
	e, err := gst.NewElementWithName("rtpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSink{Bin: &gst.Bin{Element: e}}, nil
}

// Address to send packets to (can be IPv4 or IPv6).
// Default: 0.0.0.0
func (e *GstRtpSink) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstRtpSink) GetAddress() (string, error) {
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

// The network interface on which to join the multicast group.This allows multiple interfaces separated by comma. ("eth0,eth1")
// Default: NULL
func (e *GstRtpSink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstRtpSink) GetMulticastIface() (string, error) {
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

// The port RTP packets will be sent, the RTCP port is this value + 1. This port must be an even number.
// Default: 5004, Min: 2, Max: 65534
func (e *GstRtpSink) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstRtpSink) GetPort() (uint, error) {
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

// Used for setting the unicast TTL parameter
// Default: 64, Min: 0, Max: 255
func (e *GstRtpSink) SetTtl(value int) error {
	return e.Element.SetProperty("ttl", value)
}

func (e *GstRtpSink) GetTtl() (int, error) {
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
func (e *GstRtpSink) SetTtlMc(value int) error {
	return e.Element.SetProperty("ttl-mc", value)
}

func (e *GstRtpSink) GetTtlMc() (int, error) {
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

// URI in the form of rtp://host:port?query
// Default: rtp://0.0.0.0:5004
func (e *GstRtpSink) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstRtpSink) GetUri() (string, error) {
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

// Simple RTP src
type GstRtpSrc struct {
	*gst.Bin
}

func NewRtpSrc() (*GstRtpSrc, error) {
	e, err := gst.NewElement("rtpsrc")
	if err != nil {
		return nil, err
	}
	return &GstRtpSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewRtpSrcWithName(name string) (*GstRtpSrc, error) {
	e, err := gst.NewElementWithName("rtpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSrc{Bin: &gst.Bin{Element: e}}, nil
}

// URI in the form of rtp://host:port?query
// Default: rtp://0.0.0.0:5004
func (e *GstRtpSrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstRtpSrc) GetUri() (string, error) {
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

// Address to receive packets from (can be IPv4 or IPv6).
// Default: 0.0.0.0
func (e *GstRtpSrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstRtpSrc) GetAddress() (string, error) {
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

// Default amount of ms to buffer in the jitterbuffers
// Default: 200, Min: 0, Max: -1
func (e *GstRtpSrc) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstRtpSrc) GetLatency() (uint, error) {
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

// The network interface on which to join the multicast group.This allows multiple interfaces separated by comma. ("eth0,eth1")
// Default: NULL
func (e *GstRtpSrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstRtpSrc) GetMulticastIface() (string, error) {
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

// Used for setting the unicast TTL parameter
// Default: 64, Min: 0, Max: 255
func (e *GstRtpSrc) SetTtl(value int) error {
	return e.Element.SetProperty("ttl", value)
}

func (e *GstRtpSrc) GetTtl() (int, error) {
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
func (e *GstRtpSrc) SetTtlMc(value int) error {
	return e.Element.SetProperty("ttl-mc", value)
}

func (e *GstRtpSrc) GetTtlMc() (int, error) {
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

// The caps of the incoming stream

func (e *GstRtpSrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstRtpSrc) GetCaps() (*gst.Caps, error) {
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

// Encoding name use to determine caps parameters
// Default: NULL
func (e *GstRtpSrc) SetEncodingName(value string) error {
	return e.Element.SetProperty("encoding-name", value)
}

func (e *GstRtpSrc) GetEncodingName() (string, error) {
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

// The port to listen for RTP packets, the RTCP port is this value + 1. This port must be an even number.
// Default: 5004, Min: 2, Max: 65534
func (e *GstRtpSrc) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstRtpSrc) GetPort() (uint, error) {
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
