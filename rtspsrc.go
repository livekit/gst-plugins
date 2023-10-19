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

type Rtspsrc struct {
	Element *gst.Element
}

func NewRtspsrc() (*Rtspsrc, error) {
	e, err := gst.NewElement("rtspsrc")
	if err != nil {
		return nil, err
	}
	return &Rtspsrc{Element: e}, nil
}

func NewRtspsrcWithName(name string) (*Rtspsrc, error) {
	e, err := gst.NewElementWithName("rtspsrc", name)
	if err != nil {
		return nil, err
	}
	return &Rtspsrc{Element: e}, nil
}

// ----- Properties -----

// add-reference-timestamp-meta (bool)
//
// When syncing to a RFC7273 clock, add GstReferenceTimestampMeta
// to buffers with the original reconstructed reference clock timestamp.

func (e *Rtspsrc) GetAddReferenceTimestampMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-reference-timestamp-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetAddReferenceTimestampMeta(value bool) error {
	return e.Element.SetProperty("add-reference-timestamp-meta", value)
}

// backchannel (GstRtspbackchannel)
//
// Select a type of backchannel to setup with the RTSP server.
// Default value is "none". Allowed values are "none" and "onvif".

func (e *Rtspsrc) GetBackchannel() (GstRtspbackchannel, error) {
	var value GstRtspbackchannel
	var ok bool
	v, err := e.Element.GetProperty("backchannel")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtspbackchannel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtspbackchannel")
	}
	return value, nil
}

func (e *Rtspsrc) SetBackchannel(value GstRtspbackchannel) error {
	e.Element.SetArg("backchannel", string(value))
	return nil
}

// buffer-mode (GstRtspsrcBufferMode)
//
// Control the buffering and timestamping mode used by the jitterbuffer.

func (e *Rtspsrc) GetBufferMode() (GstRtspsrcBufferMode, error) {
	var value GstRtspsrcBufferMode
	var ok bool
	v, err := e.Element.GetProperty("buffer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtspsrcBufferMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtspsrcBufferMode")
	}
	return value, nil
}

func (e *Rtspsrc) SetBufferMode(value GstRtspsrcBufferMode) error {
	e.Element.SetArg("buffer-mode", string(value))
	return nil
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Rtspsrc) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtspsrc) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// debug (bool)
//
// Dump request and response messages to stdout(DEPRECATED: Printed all RTSP message to gstreamer log as 'log' level)

func (e *Rtspsrc) GetDebug() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("debug")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetDebug(value bool) error {
	return e.Element.SetProperty("debug", value)
}

// default-rtsp-version (GstRTSPVersion)
//
// The preferred RTSP version to use while negotiating the version with the server.

func (e *Rtspsrc) GetDefaultRtspVersion() (interface{}, error) {
	return e.Element.GetProperty("default-rtsp-version")
}

func (e *Rtspsrc) SetDefaultRtspVersion(value interface{}) error {
	return e.Element.SetProperty("default-rtsp-version", value)
}

// do-retransmission (bool)
//
// Ask the server to retransmit lost packets

func (e *Rtspsrc) GetDoRetransmission() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-retransmission")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetDoRetransmission(value bool) error {
	return e.Element.SetProperty("do-retransmission", value)
}

// do-rtcp (bool)
//
// Enable RTCP support. Some old server don't like RTCP and then this property
// needs to be set to FALSE.

func (e *Rtspsrc) GetDoRtcp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-rtcp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetDoRtcp(value bool) error {
	return e.Element.SetProperty("do-rtcp", value)
}

// do-rtsp-keep-alive (bool)
//
// Enable RTSP keep alive support. Some old server don't like RTSP
// keep alive and then this property needs to be set to FALSE.

func (e *Rtspsrc) GetDoRtspKeepAlive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-rtsp-keep-alive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetDoRtspKeepAlive(value bool) error {
	return e.Element.SetProperty("do-rtsp-keep-alive", value)
}

// drop-on-latency (bool)
//
// Tells the jitterbuffer to never exceed the given latency in size

func (e *Rtspsrc) GetDropOnLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-on-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetDropOnLatency(value bool) error {
	return e.Element.SetProperty("drop-on-latency", value)
}

// extra-http-request-headers (GstStructure)
//
// When in tunneled mode append provided headers to any HTTP requests
// made by rtspsrc.

// ignore-x-server-reply (bool)
//
// When connecting to an RTSP server in tunneled mode (HTTP) the server
// usually replies with an x-server-ip-address header. This contains the
// address of the intended streaming server. However some servers return an
// "invalid" address. Here follows two examples when it might happen.

// is-live (bool)
//
// Whether to act as a live source. This is useful in combination with
// onvif-rate-control set to FALSE and usage of the TCP
// protocol. In that situation, data delivery rate can be entirely
// controlled from the client side, enabling features such as frame
// stepping and instantaneous rate changes.

func (e *Rtspsrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// latency (uint)
//
// Amount of ms to buffer

func (e *Rtspsrc) GetLatency() (uint, error) {
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

func (e *Rtspsrc) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

// location (string)
//
// Location of the RTSP url to read

func (e *Rtspsrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-rtcp-rtp-time-diff (int)
//
// Maximum amount of time in ms that the RTP time in RTCP SRs is allowed to be ahead (-1 disabled)

func (e *Rtspsrc) GetMaxRtcpRtpTimeDiff() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-rtcp-rtp-time-diff")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtspsrc) SetMaxRtcpRtpTimeDiff(value int) error {
	return e.Element.SetProperty("max-rtcp-rtp-time-diff", value)
}

// max-ts-offset (int64)
//
// Used to set an upper limit of how large a time offset may be. This
// is used to protect against unrealistic values as a result of either
// client,server or clock issues.

func (e *Rtspsrc) GetMaxTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Rtspsrc) SetMaxTsOffset(value int64) error {
	return e.Element.SetProperty("max-ts-offset", value)
}

// max-ts-offset-adjustment (uint64)
//
// Syncing time stamps to NTP time adds a time offset. This parameter
// specifies the maximum number of nanoseconds per frame that this time offset
// may be adjusted with. This is used to avoid sudden large changes to time
// stamps.

func (e *Rtspsrc) GetMaxTsOffsetAdjustment() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-ts-offset-adjustment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtspsrc) SetMaxTsOffsetAdjustment(value uint64) error {
	return e.Element.SetProperty("max-ts-offset-adjustment", value)
}

// multicast-iface (string)
//
// The network interface on which to join the multicast group

func (e *Rtspsrc) GetMulticastIface() (string, error) {
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

func (e *Rtspsrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// nat-method (GstRtspnatMethod)
//
// Method to use for traversing firewalls and NAT

func (e *Rtspsrc) GetNatMethod() (GstRtspnatMethod, error) {
	var value GstRtspnatMethod
	var ok bool
	v, err := e.Element.GetProperty("nat-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtspnatMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtspnatMethod")
	}
	return value, nil
}

func (e *Rtspsrc) SetNatMethod(value GstRtspnatMethod) error {
	e.Element.SetArg("nat-method", string(value))
	return nil
}

// ntp-sync (bool)
//
// Synchronize received streams to the NTP clock

func (e *Rtspsrc) GetNtpSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ntp-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetNtpSync(value bool) error {
	return e.Element.SetProperty("ntp-sync", value)
}

// ntp-time-source (GstRtspsrcNtpTimeSource)
//
// NTP time source for RTCP packets

func (e *Rtspsrc) GetNtpTimeSource() (GstRtspsrcNtpTimeSource, error) {
	var value GstRtspsrcNtpTimeSource
	var ok bool
	v, err := e.Element.GetProperty("ntp-time-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtspsrcNtpTimeSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtspsrcNtpTimeSource")
	}
	return value, nil
}

func (e *Rtspsrc) SetNtpTimeSource(value GstRtspsrcNtpTimeSource) error {
	e.Element.SetArg("ntp-time-source", string(value))
	return nil
}

// onvif-mode (bool)
//
// Act as an ONVIF client. When set to TRUE:

// onvif-rate-control (bool)
//
// When in onvif-mode, whether to set Rate-Control to yes or no. When set
// to FALSE, the server will deliver data as fast as the client can consume
// it.

func (e *Rtspsrc) GetOnvifRateControl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("onvif-rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetOnvifRateControl(value bool) error {
	return e.Element.SetProperty("onvif-rate-control", value)
}

// port-range (string)
//
// Configure the client port numbers that can be used to receive RTP and
// RTCP.

func (e *Rtspsrc) GetPortRange() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("port-range")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetPortRange(value string) error {
	return e.Element.SetProperty("port-range", value)
}

// probation (uint)
//
// Consecutive packet sequence numbers to accept the source

func (e *Rtspsrc) GetProbation() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("probation")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtspsrc) SetProbation(value uint) error {
	return e.Element.SetProperty("probation", value)
}

// protocols (GstRTSPLowerTrans)
//
// Allowed lower transport protocols

func (e *Rtspsrc) GetProtocols() (interface{}, error) {
	return e.Element.GetProperty("protocols")
}

func (e *Rtspsrc) SetProtocols(value interface{}) error {
	return e.Element.SetProperty("protocols", value)
}

// proxy (string)
//
// Set the proxy parameters. This has to be a string of the format
// [http://][user:passwd@]host[:port].

func (e *Rtspsrc) GetProxy() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// proxy-id (string)
//
// Sets the proxy URI user id for authentication. If the URI set via the
// "proxy" property contains a user-id already, that will take precedence.

func (e *Rtspsrc) GetProxyId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

// proxy-pw (string)
//
// Sets the proxy URI password for authentication. If the URI set via the
// "proxy" property contains a password already, that will take precedence.

func (e *Rtspsrc) GetProxyPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

// retry (uint)
//
// Max number of retries when allocating RTP ports.

func (e *Rtspsrc) GetRetry() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("retry")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtspsrc) SetRetry(value uint) error {
	return e.Element.SetProperty("retry", value)
}

// rfc7273-sync (bool)
//
// Synchronize received streams to the RFC7273 clock (requires clock and offset to be provided)

func (e *Rtspsrc) GetRfc7273Sync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rfc7273-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetRfc7273Sync(value bool) error {
	return e.Element.SetProperty("rfc7273-sync", value)
}

// rtp-blocksize (uint)
//
// RTP package size to suggest to server.

func (e *Rtspsrc) GetRtpBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtp-blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtspsrc) SetRtpBlocksize(value uint) error {
	return e.Element.SetProperty("rtp-blocksize", value)
}

// sdes (GstStructure)
//
// The SDES items of this session

func (e *Rtspsrc) GetSdes() (interface{}, error) {
	return e.Element.GetProperty("sdes")
}

func (e *Rtspsrc) SetSdes(value interface{}) error {
	return e.Element.SetProperty("sdes", value)
}

// short-header (bool)
//
// Only send the basic RTSP headers for broken encoders.

func (e *Rtspsrc) GetShortHeader() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("short-header")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetShortHeader(value bool) error {
	return e.Element.SetProperty("short-header", value)
}

// tcp-timeout (uint64)
//
// Fail after timeout microseconds on TCP connections (0 = disabled)

func (e *Rtspsrc) GetTcpTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("tcp-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtspsrc) SetTcpTimeout(value uint64) error {
	return e.Element.SetProperty("tcp-timeout", value)
}

// teardown-timeout (uint64)
//
// When transitioning PAUSED-READY, allow up to timeout (in nanoseconds)
// delay in order to send teardown (0 = disabled)

func (e *Rtspsrc) GetTeardownTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("teardown-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtspsrc) SetTeardownTimeout(value uint64) error {
	return e.Element.SetProperty("teardown-timeout", value)
}

// timeout (uint64)
//
// Retry TCP transport after UDP timeout microseconds (0 = disabled)

func (e *Rtspsrc) GetTimeout() (uint64, error) {
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

func (e *Rtspsrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

// tls-database (GstGtlsDatabase)
//
// TLS database with anchor certificate authorities used to validate the server certificate

func (e *Rtspsrc) GetTlsDatabase() (interface{}, error) {
	return e.Element.GetProperty("tls-database")
}

func (e *Rtspsrc) SetTlsDatabase(value interface{}) error {
	return e.Element.SetProperty("tls-database", value)
}

// tls-interaction (GstGtlsInteraction)
//
// A GTlsInteraction object to prompt the user for password or certificate

func (e *Rtspsrc) GetTlsInteraction() (interface{}, error) {
	return e.Element.GetProperty("tls-interaction")
}

func (e *Rtspsrc) SetTlsInteraction(value interface{}) error {
	return e.Element.SetProperty("tls-interaction", value)
}

// tls-validation-flags (GstGtlsCertificateFlags)
//
// TLS certificate validation flags used to validate the server certificate

func (e *Rtspsrc) GetTlsValidationFlags() (interface{}, error) {
	return e.Element.GetProperty("tls-validation-flags")
}

func (e *Rtspsrc) SetTlsValidationFlags(value interface{}) error {
	return e.Element.SetProperty("tls-validation-flags", value)
}

// udp-buffer-size (int)
//
// Size of the kernel UDP receive buffer in bytes.

func (e *Rtspsrc) GetUdpBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("udp-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtspsrc) SetUdpBufferSize(value int) error {
	return e.Element.SetProperty("udp-buffer-size", value)
}

// udp-reconnect (bool)
//
// Reconnect to the server if RTSP connection is closed when doing UDP

func (e *Rtspsrc) GetUdpReconnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("udp-reconnect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetUdpReconnect(value bool) error {
	return e.Element.SetProperty("udp-reconnect", value)
}

// use-pipeline-clock (bool)
//
// Use the pipeline running-time to set the NTP time in the RTCP SR messages(DEPRECATED: Use ntp-time-source property)

func (e *Rtspsrc) GetUsePipelineClock() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-pipeline-clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtspsrc) SetUsePipelineClock(value bool) error {
	return e.Element.SetProperty("use-pipeline-clock", value)
}

// user-agent (string)
//
// The User-Agent string to send to the server

func (e *Rtspsrc) GetUserAgent() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-agent")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

// user-id (string)
//
// RTSP location URI user id for authentication

func (e *Rtspsrc) GetUserId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

// user-pw (string)
//
// RTSP location URI user password for authentication

func (e *Rtspsrc) GetUserPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtspsrc) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

// ----- Constants -----

type GstRtspbackchannel string

const (
	GstRtspbackchannelNone GstRtspbackchannel = "none" // none (0) – No backchannel
	GstRtspbackchannelOnvif GstRtspbackchannel = "onvif" // onvif (1) – ONVIF audio backchannel
)

type GstRtspnatMethod string

const (
	GstRtspnatMethodNone GstRtspnatMethod = "none" // none (0) – None
	GstRtspnatMethodDummy GstRtspnatMethod = "dummy" // dummy (1) – Send Dummy packets
)

type GstRtspsrcBufferMode string

const (
	GstRtspsrcBufferModeNone GstRtspsrcBufferMode = "none" // none (0) – Only use RTP timestamps
	GstRtspsrcBufferModeSlave GstRtspsrcBufferMode = "slave" // slave (1) – Slave receiver to sender clock
	GstRtspsrcBufferModeBuffer GstRtspsrcBufferMode = "buffer" // buffer (2) – Do low/high watermark buffering
	GstRtspsrcBufferModeAuto GstRtspsrcBufferMode = "auto" // auto (3) – Choose mode depending on stream live
	GstRtspsrcBufferModeSynced GstRtspsrcBufferMode = "synced" // synced (4) – Synchronized sender and receiver clocks
)

type GstRtspsrcNtpTimeSource string

const (
	GstRtspsrcNtpTimeSourceNtp GstRtspsrcNtpTimeSource = "ntp" // ntp (0) – NTP time based on realtime clock
	GstRtspsrcNtpTimeSourceUnix GstRtspsrcNtpTimeSource = "unix" // unix (1) – UNIX time based on realtime clock
	GstRtspsrcNtpTimeSourceRunningTime GstRtspsrcNtpTimeSource = "running-time" // running-time (2) – Running time based on pipeline clock
	GstRtspsrcNtpTimeSourceClockTime GstRtspsrcNtpTimeSource = "clock-time" // clock-time (3) – Pipeline clock time
)

