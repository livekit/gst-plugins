// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// A round robin dispatcher element.
type GstRoundRobin struct {
	*gst.Element
}

func NewRoundRobin() (*GstRoundRobin, error) {
	e, err := gst.NewElement("roundrobin")
	if err != nil {
		return nil, err
	}
	return &GstRoundRobin{Element: e}, nil
}

func NewRoundRobinWithName(name string) (*GstRoundRobin, error) {
	e, err := gst.NewElementWithName("roundrobin", name)
	if err != nil {
		return nil, err
	}
	return &GstRoundRobin{Element: e}, nil
}

// Removes RIST TR-06-2 RTP Header extension
type GstRistRtpDeext struct {
	*gst.Element
}

func NewRistRtpDeext() (*GstRistRtpDeext, error) {
	e, err := gst.NewElement("ristrtpdeext")
	if err != nil {
		return nil, err
	}
	return &GstRistRtpDeext{Element: e}, nil
}

func NewRistRtpDeextWithName(name string) (*GstRistRtpDeext, error) {
	e, err := gst.NewElementWithName("ristrtpdeext", name)
	if err != nil {
		return nil, err
	}
	return &GstRistRtpDeext{Element: e}, nil
}

// Largest extended sequence number received
// Default: 0, Min: 0, Max: -1
func (e *GstRistRtpDeext) GetMaxExtSeqnum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-ext-seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Has an extended sequence number extension been seen
// Default: false
func (e *GstRistRtpDeext) GetHaveExtSeqnum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("have-ext-seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adds RIST TR-06-2 RTP Header extension
type GstRistRtpExt struct {
	*gst.Element
}

func NewRistRtpExt() (*GstRistRtpExt, error) {
	e, err := gst.NewElement("ristrtpext")
	if err != nil {
		return nil, err
	}
	return &GstRistRtpExt{Element: e}, nil
}

func NewRistRtpExtWithName(name string) (*GstRistRtpExt, error) {
	e, err := gst.NewElementWithName("ristrtpext", name)
	if err != nil {
		return nil, err
	}
	return &GstRistRtpExt{Element: e}, nil
}

// Drop null MPEG-TS packet and replace them with a custom header extension.
// Default: false
func (e *GstRistRtpExt) SetDropNullTsPackets(value bool) error {
	return e.Element.SetProperty("drop-null-ts-packets", value)
}

func (e *GstRistRtpExt) GetDropNullTsPackets() (bool, error) {
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

// Add sequence number extension to packets.
// Default: false
func (e *GstRistRtpExt) SetSequenceNumberExtension(value bool) error {
	return e.Element.SetProperty("sequence-number-extension", value)
}

func (e *GstRistRtpExt) GetSequenceNumberExtension() (bool, error) {
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

// Receive retransmitted RIST packets according to VSF TR-06-1
type GstRistRtxReceive struct {
	*gst.Element
}

func NewRistRtxReceive() (*GstRistRtxReceive, error) {
	e, err := gst.NewElement("ristrtxreceive")
	if err != nil {
		return nil, err
	}
	return &GstRistRtxReceive{Element: e}, nil
}

func NewRistRtxReceiveWithName(name string) (*GstRistRtxReceive, error) {
	e, err := gst.NewElementWithName("ristrtxreceive", name)
	if err != nil {
		return nil, err
	}
	return &GstRistRtxReceive{Element: e}, nil
}

//	Number of retransmission packets received
//
// Default: 0, Min: 0, Max: -1
func (e *GstRistRtxReceive) GetNumRtxPackets() (uint, error) {
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

// Number of retransmission events received
// Default: 0, Min: 0, Max: -1
func (e *GstRistRtxReceive) GetNumRtxRequests() (uint, error) {
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

// Retransmit RTP packets when needed, according to VSF TR-06-1
type GstRistRtxSend struct {
	*gst.Element
}

func NewRistRtxSend() (*GstRistRtxSend, error) {
	e, err := gst.NewElement("ristrtxsend")
	if err != nil {
		return nil, err
	}
	return &GstRistRtxSend{Element: e}, nil
}

func NewRistRtxSendWithName(name string) (*GstRistRtxSend, error) {
	e, err := gst.NewElementWithName("ristrtxsend", name)
	if err != nil {
		return nil, err
	}
	return &GstRistRtxSend{Element: e}, nil
}

// Amount of packets to queue (0 = unlimited)
// Default: 100, Min: 0, Max: 32767
func (e *GstRistRtxSend) SetMaxSizePackets(value uint) error {
	return e.Element.SetProperty("max-size-packets", value)
}

func (e *GstRistRtxSend) GetMaxSizePackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Amount of ms to queue (0 = unlimited)
// Default: 0, Min: 0, Max: -1
func (e *GstRistRtxSend) SetMaxSizeTime(value uint) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstRistRtxSend) GetMaxSizeTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

//	Number of retransmission packets sent
//
// Default: 0, Min: 0, Max: -1
func (e *GstRistRtxSend) GetNumRtxPackets() (uint, error) {
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

// Number of retransmission events received
// Default: 0, Min: 0, Max: -1
func (e *GstRistRtxSend) GetNumRtxRequests() (uint, error) {
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

// Sink that implements RIST TR-06-1 streaming specification
type GstRistSink struct {
	*gst.Bin
}

func NewRistSink() (*GstRistSink, error) {
	e, err := gst.NewElement("ristsink")
	if err != nil {
		return nil, err
	}
	return &GstRistSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewRistSinkWithName(name string) (*GstRistSink, error) {
	e, err := gst.NewElementWithName("ristsink", name)
	if err != nil {
		return nil, err
	}
	return &GstRistSink{Bin: &gst.Bin{Element: e}}, nil
}

// Address to send packets to (can be IPv4 or IPv6).
// Default: localhost
func (e *GstRistSink) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstRistSink) GetAddress() (string, error) {
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

// The maximum bandwidth used for RTCP as a fraction of RTP bandwdith
// Default: 0.05, Min: 0, Max: 0.05
func (e *GstRistSink) SetMaxRtcpBandwidth(value float64) error {
	return e.Element.SetProperty("max-rtcp-bandwidth", value)
}

func (e *GstRistSink) GetMaxRtcpBandwidth() (float64, error) {
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

// The minimum interval (in ms) between two regular successive RTCP packets.
// Default: 100, Min: 0, Max: 100
func (e *GstRistSink) SetMinRtcpInterval(value uint) error {
	return e.Element.SetProperty("min-rtcp-interval", value)
}

func (e *GstRistSink) GetMinRtcpInterval() (uint, error) {
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

// Statistic in a GstStructure named 'rist/x-sender-stats'
// Default: rist/x-sender-stats, sent-original-packets=(guint64)0, sent-retransmitted-packets=(guint64)0, session-stats=(GValueArray)< "rist/x-sender-session-stats\,\ session-id\=\(int\)0\,\ sent-original-packets\=\(guint64\)0\,\ sent-retransmitted-packets\=\(guint64\)0\,\ round-trip-time\=\(guint64\)0\;" >;
func (e *GstRistSink) GetStats() (*gst.Structure, error) {
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

// Comma (,) separated list of <address>:<port> to send to.
// Default: localhost:5004
func (e *GstRistSink) SetBondingAddresses(value string) error {
	return e.Element.SetProperty("bonding-addresses", value)
}

func (e *GstRistSink) GetBondingAddresses() (string, error) {
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

// Set the CNAME in the SDES block of the sender report.
// Default: NULL
func (e *GstRistSink) SetCname(value string) error {
	return e.Element.SetProperty("cname", value)
}

func (e *GstRistSink) GetCname() (string, error) {
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

// An element that takes care of multi-plexing bonded links. When set "bonding-method" is ignored.

func (e *GstRistSink) SetDispatcher(value *gst.Element) error {
	return e.Element.SetProperty("dispatcher", value)
}

func (e *GstRistSink) GetDispatcher() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("dispatcher")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// When enabled, the packet will be received locally.
// Default: false
func (e *GstRistSink) SetMulticastLoopback(value bool) error {
	return e.Element.SetProperty("multicast-loopback", value)
}

func (e *GstRistSink) GetMulticastLoopback() (bool, error) {
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

// The port RTP packets will be sent, the RTCP port is this value + 1. This port must be an even number.
// Default: 5004, Min: 2, Max: 65534
func (e *GstRistSink) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstRistSink) GetPort() (uint, error) {
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

// Size of the retransmission queue (in ms)
// Default: 1200, Min: 0, Max: -1
func (e *GstRistSink) SetSenderBuffer(value uint) error {
	return e.Element.SetProperty("sender-buffer", value)
}

func (e *GstRistSink) GetSenderBuffer() (uint, error) {
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

// Add sequence number extension to packets.
// Default: false
func (e *GstRistSink) SetSequenceNumberExtension(value bool) error {
	return e.Element.SetProperty("sequence-number-extension", value)
}

func (e *GstRistSink) GetSequenceNumberExtension() (bool, error) {
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

// Defines the bonding method to use.
// Default: broadcast (0)
func (e *GstRistSink) SetBondingMethod(value GstRistBondingMethodType) error {
	e.Element.SetArg("bonding-method", string(value))
	return nil
}

func (e *GstRistSink) GetBondingMethod() (GstRistBondingMethodType, error) {
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

// Drop null MPEG-TS packet and replace them with a custom header extension.
// Default: false
func (e *GstRistSink) SetDropNullTsPackets(value bool) error {
	return e.Element.SetProperty("drop-null-ts-packets", value)
}

func (e *GstRistSink) GetDropNullTsPackets() (bool, error) {
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

// The multicast interface to use to send packets.
// Default: NULL
func (e *GstRistSink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstRistSink) GetMulticastIface() (string, error) {
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

// The multicast time-to-live parameter.
// Default: 1, Min: 0, Max: 255
func (e *GstRistSink) SetMulticastTtl(value int) error {
	return e.Element.SetProperty("multicast-ttl", value)
}

func (e *GstRistSink) GetMulticastTtl() (int, error) {
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

// The interval between 'stats' update notification (in ms) (0 disabled)
// Default: 0, Min: 0, Max: -1
func (e *GstRistSink) SetStatsUpdateInterval(value uint) error {
	return e.Element.SetProperty("stats-update-interval", value)
}

func (e *GstRistSink) GetStatsUpdateInterval() (uint, error) {
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

// Source that implements RIST TR-06-1 streaming specification
type GstRistSrc struct {
	*gst.Bin
}

func NewRistSrc() (*GstRistSrc, error) {
	e, err := gst.NewElement("ristsrc")
	if err != nil {
		return nil, err
	}
	return &GstRistSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewRistSrcWithName(name string) (*GstRistSrc, error) {
	e, err := gst.NewElementWithName("ristsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRistSrc{Bin: &gst.Bin{Element: e}}, nil
}

// Statistic in a GstStructure named 'rist/x-receiver-stats'
// Default: rist/x-receiver-stats, dropped=(guint64)0, received=(guint64)0, recovered=(guint64)0, permanently-lost=(guint64)0, duplicates=(guint64)0, retransmission-requests-sent=(guint64)0, rtx-roundtrip-time=(guint64)0, session-stats=(GValueArray)< "rist/x-receiver-session-stats\,\ session-id\=\(int\)0\,\ rtp-from\=\(string\)\"\"\,\ rtcp-from\=\(string\)\"\"\,\ dropped\=\(guint64\)0\,\ received\=\(guint64\)0\;" >;
func (e *GstRistSrc) GetStats() (*gst.Structure, error) {
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

// Comma (,) separated list of <address>:<port> to receive from. Only used if 'enable-bonding' is set.
// Default: 0.0.0.0:5004
func (e *GstRistSrc) SetBondingAddresses(value string) error {
	return e.Element.SetProperty("bonding-addresses", value)
}

func (e *GstRistSrc) GetBondingAddresses() (string, error) {
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

// The multicast interface to use to send packets.
// Default: NULL
func (e *GstRistSrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstRistSrc) GetMulticastIface() (string, error) {
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

// The port to listen for RTP packets, the RTCP port is this value + 1. This port must be an even number.
// Default: 5004, Min: 2, Max: 65534
func (e *GstRistSrc) SetPort(value uint) error {
	return e.Element.SetProperty("port", value)
}

func (e *GstRistSrc) GetPort() (uint, error) {
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

// The caps of the incoming stream

func (e *GstRistSrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstRistSrc) GetCaps() (*gst.Caps, error) {
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

// The multicast time-to-live parameter.
// Default: 1, Min: 0, Max: 255
func (e *GstRistSrc) SetMulticastTtl(value int) error {
	return e.Element.SetProperty("multicast-ttl", value)
}

func (e *GstRistSrc) GetMulticastTtl() (int, error) {
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

// When enabled, the packets will be received locally.
// Default: false
func (e *GstRistSrc) SetMulticastLoopback(value bool) error {
	return e.Element.SetProperty("multicast-loopback", value)
}

func (e *GstRistSrc) GetMulticastLoopback() (bool, error) {
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

// Buffering duration (in ms)
// Default: 1000, Min: 0, Max: -1
func (e *GstRistSrc) SetReceiverBuffer(value uint) error {
	return e.Element.SetProperty("receiver-buffer", value)
}

func (e *GstRistSrc) GetReceiverBuffer() (uint, error) {
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

// The interval between 'stats' update notification (in ms) (0 disabled)
// Default: 0, Min: 0, Max: -1
func (e *GstRistSrc) SetStatsUpdateInterval(value uint) error {
	return e.Element.SetProperty("stats-update-interval", value)
}

func (e *GstRistSrc) GetStatsUpdateInterval() (uint, error) {
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

// The maximum bandwidth used for RTCP as a fraction of RTP bandwdith
// Default: 0.05, Min: 0, Max: 0.05
func (e *GstRistSrc) SetMaxRtcpBandwidth(value float64) error {
	return e.Element.SetProperty("max-rtcp-bandwidth", value)
}

func (e *GstRistSrc) GetMaxRtcpBandwidth() (float64, error) {
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

// The maximum number of retransmission requests for a lost packet.
// Default: 7, Min: 0, Max: -1
func (e *GstRistSrc) SetMaxRtxRetries(value uint) error {
	return e.Element.SetProperty("max-rtx-retries", value)
}

func (e *GstRistSrc) GetMaxRtxRetries() (uint, error) {
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

// The minimum interval (in ms) between two successive RTCP packets
// Default: 100, Min: 0, Max: 100
func (e *GstRistSrc) SetMinRtcpInterval(value uint) error {
	return e.Element.SetProperty("min-rtcp-interval", value)
}

func (e *GstRistSrc) GetMinRtcpInterval() (uint, error) {
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

// Time to wait before sending retransmission request (in ms)
// Default: 70, Min: 0, Max: -1
func (e *GstRistSrc) SetReorderSection(value uint) error {
	return e.Element.SetProperty("reorder-section", value)
}

func (e *GstRistSrc) GetReorderSection() (uint, error) {
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

// Address to receive packets from (can be IPv4 or IPv6).
// Default: 0.0.0.0
func (e *GstRistSrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstRistSrc) GetAddress() (string, error) {
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

// Set the CNAME in the SDES block of the receiver report.
// Default: NULL
func (e *GstRistSrc) SetCname(value string) error {
	return e.Element.SetProperty("cname", value)
}

func (e *GstRistSrc) GetCname() (string, error) {
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

// Encoding name use to determine caps parameters
// Default: NULL
func (e *GstRistSrc) SetEncodingName(value string) error {
	return e.Element.SetProperty("encoding-name", value)
}

func (e *GstRistSrc) GetEncodingName() (string, error) {
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

type GstRistBondingMethodType string

const (
	GstRistBondingMethodTypeBroadcast  GstRistBondingMethodType = "broadcast"   // GST_RIST_BONDING_METHOD_BROADCAST
	GstRistBondingMethodTypeRoundRobin GstRistBondingMethodType = "round-robin" // GST_RIST_BONDING_METHOD_ROUND_ROBIN
)
