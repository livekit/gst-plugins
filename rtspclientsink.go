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

type Rtspclientsink struct {
	Element *gst.Element
}

func NewRtspclientsink() (*Rtspclientsink, error) {
	e, err := gst.NewElement("rtspclientsink")
	if err != nil {
		return nil, err
	}
	return &Rtspclientsink{Element: e}, nil
}

func NewRtspclientsinkWithName(name string) (*Rtspclientsink, error) {
	e, err := gst.NewElementWithName("rtspclientsink", name)
	if err != nil {
		return nil, err
	}
	return &Rtspclientsink{Element: e}, nil
}

// ----- Properties -----

// debug (bool)
//
// Dump request and response messages to stdout

func (e *Rtspclientsink) GetDebug() (bool, error) {
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

func (e *Rtspclientsink) SetDebug(value bool) error {
	return e.Element.SetProperty("debug", value)
}

// do-rtsp-keep-alive (bool)
//
// Enable RTSP keep alive support. Some old server don't like RTSP
// keep alive and then this property needs to be set to FALSE.

func (e *Rtspclientsink) GetDoRtspKeepAlive() (bool, error) {
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

func (e *Rtspclientsink) SetDoRtspKeepAlive(value bool) error {
	return e.Element.SetProperty("do-rtsp-keep-alive", value)
}

// latency (uint)
//
// Amount of ms to buffer

func (e *Rtspclientsink) GetLatency() (uint, error) {
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

func (e *Rtspclientsink) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

// location (string)
//
// Location of the RTSP url to read

func (e *Rtspclientsink) GetLocation() (string, error) {
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

func (e *Rtspclientsink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// multicast-iface (string)
//
// The network interface on which to join the multicast group

func (e *Rtspclientsink) GetMulticastIface() (string, error) {
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

func (e *Rtspclientsink) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

// ntp-time-source (GstRTSPClientSinkNtpTimeSource)
//
// allows to select the time source that should be used
// for the NTP time in outgoing packets

func (e *Rtspclientsink) GetNtpTimeSource() (interface{}, error) {
	return e.Element.GetProperty("ntp-time-source")
}

func (e *Rtspclientsink) SetNtpTimeSource(value interface{}) error {
	return e.Element.SetProperty("ntp-time-source", value)
}

// port-range (string)
//
// Configure the client port numbers that can be used to send RTP and receive
// RTCP.

func (e *Rtspclientsink) GetPortRange() (string, error) {
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

func (e *Rtspclientsink) SetPortRange(value string) error {
	return e.Element.SetProperty("port-range", value)
}

// profiles (GstRTSPProfile)
//
// Allowed RTSP profiles

func (e *Rtspclientsink) GetProfiles() (interface{}, error) {
	return e.Element.GetProperty("profiles")
}

func (e *Rtspclientsink) SetProfiles(value interface{}) error {
	return e.Element.SetProperty("profiles", value)
}

// protocols (GstRTSPLowerTrans)
//
// Allowed lower transport protocols

func (e *Rtspclientsink) GetProtocols() (interface{}, error) {
	return e.Element.GetProperty("protocols")
}

func (e *Rtspclientsink) SetProtocols(value interface{}) error {
	return e.Element.SetProperty("protocols", value)
}

// proxy (string)
//
// Set the proxy parameters. This has to be a string of the format
// [http://][user:passwd@]host[:port].

func (e *Rtspclientsink) GetProxy() (string, error) {
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

func (e *Rtspclientsink) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

// proxy-id (string)
//
// Sets the proxy URI user id for authentication. If the URI set via the
// "proxy" property contains a user-id already, that will take precedence.

func (e *Rtspclientsink) GetProxyId() (string, error) {
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

func (e *Rtspclientsink) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

// proxy-pw (string)
//
// Sets the proxy URI password for authentication. If the URI set via the
// "proxy" property contains a password already, that will take precedence.

func (e *Rtspclientsink) GetProxyPw() (string, error) {
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

func (e *Rtspclientsink) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

// publish-clock-mode (GstRTSPPublishClockMode)
//
// Sets if and how the media clock should be published according to RFC7273.

func (e *Rtspclientsink) GetPublishClockMode() (interface{}, error) {
	return e.Element.GetProperty("publish-clock-mode")
}

func (e *Rtspclientsink) SetPublishClockMode(value interface{}) error {
	return e.Element.SetProperty("publish-clock-mode", value)
}

// retry (uint)
//
// Max number of retries when allocating RTP ports.

func (e *Rtspclientsink) GetRetry() (uint, error) {
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

func (e *Rtspclientsink) SetRetry(value uint) error {
	return e.Element.SetProperty("retry", value)
}

// rtp-blocksize (uint)
//
// RTP package size to suggest to server.

func (e *Rtspclientsink) GetRtpBlocksize() (uint, error) {
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

func (e *Rtspclientsink) SetRtpBlocksize(value uint) error {
	return e.Element.SetProperty("rtp-blocksize", value)
}

// rtx-time (uint)
//
// Amount of ms to buffer for retransmission. 0 disables retransmission

func (e *Rtspclientsink) GetRtxTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtx-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtspclientsink) SetRtxTime(value uint) error {
	return e.Element.SetProperty("rtx-time", value)
}

// sdes (GstStructure)
//
// The SDES items of this session

func (e *Rtspclientsink) GetSdes() (interface{}, error) {
	return e.Element.GetProperty("sdes")
}

func (e *Rtspclientsink) SetSdes(value interface{}) error {
	return e.Element.SetProperty("sdes", value)
}

// tcp-timeout (uint64)
//
// Fail after timeout microseconds on TCP connections (0 = disabled)

func (e *Rtspclientsink) GetTcpTimeout() (uint64, error) {
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

func (e *Rtspclientsink) SetTcpTimeout(value uint64) error {
	return e.Element.SetProperty("tcp-timeout", value)
}

// timeout (uint64)
//
// Retry TCP transport after UDP timeout microseconds (0 = disabled)

func (e *Rtspclientsink) GetTimeout() (uint64, error) {
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

func (e *Rtspclientsink) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

// tls-database (GstGtlsDatabase)
//
// TLS database with anchor certificate authorities used to validate
// the server certificate.

func (e *Rtspclientsink) GetTlsDatabase() (interface{}, error) {
	return e.Element.GetProperty("tls-database")
}

func (e *Rtspclientsink) SetTlsDatabase(value interface{}) error {
	return e.Element.SetProperty("tls-database", value)
}

// tls-interaction (GstGtlsInteraction)
//
// A GTlsInteraction object to be used when the connection or certificate
// database need to interact with the user. This will be used to prompt the
// user for passwords where necessary.

func (e *Rtspclientsink) GetTlsInteraction() (interface{}, error) {
	return e.Element.GetProperty("tls-interaction")
}

func (e *Rtspclientsink) SetTlsInteraction(value interface{}) error {
	return e.Element.SetProperty("tls-interaction", value)
}

// tls-validation-flags (GstGtlsCertificateFlags)
//
// TLS certificate validation flags used to validate server
// certificate.

// udp-buffer-size (int)
//
// Size of the kernel UDP receive buffer in bytes.

func (e *Rtspclientsink) GetUdpBufferSize() (int, error) {
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

func (e *Rtspclientsink) SetUdpBufferSize(value int) error {
	return e.Element.SetProperty("udp-buffer-size", value)
}

// udp-reconnect (bool)
//
// Reconnect to the server if RTSP connection is closed when doing UDP

func (e *Rtspclientsink) GetUdpReconnect() (bool, error) {
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

func (e *Rtspclientsink) SetUdpReconnect(value bool) error {
	return e.Element.SetProperty("udp-reconnect", value)
}

// user-agent (string)
//
// The string to set in the User-Agent header.

func (e *Rtspclientsink) GetUserAgent() (string, error) {
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

func (e *Rtspclientsink) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

// user-id (string)
//
// RTSP location URI user id for authentication

func (e *Rtspclientsink) GetUserId() (string, error) {
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

func (e *Rtspclientsink) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

// user-pw (string)
//
// RTSP location URI user password for authentication

func (e *Rtspclientsink) GetUserPw() (string, error) {
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

func (e *Rtspclientsink) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

// payloader (GstElement)
//
// The payloader element to use (NULL = default automatically selected)

func (e *Rtspclientsink) GetPayloader() (interface{}, error) {
	return e.Element.GetProperty("payloader")
}

func (e *Rtspclientsink) SetPayloader(value interface{}) error {
	return e.Element.SetProperty("payloader", value)
}

// ulpfec-percentage (uint)
//
// The percentage of ULP redundancy to apply

func (e *Rtspclientsink) GetUlpfecPercentage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ulpfec-percentage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtspclientsink) SetUlpfecPercentage(value uint) error {
	return e.Element.SetProperty("ulpfec-percentage", value)
}

// ----- Constants -----

type GstRtspclientSinkNtpTimeSource string

const (
	GstRtspclientSinkNtpTimeSourceNtp GstRtspclientSinkNtpTimeSource = "ntp" // ntp (0) – NTP time based on realtime clock
	GstRtspclientSinkNtpTimeSourceUnix GstRtspclientSinkNtpTimeSource = "unix" // unix (1) – UNIX time based on realtime clock
	GstRtspclientSinkNtpTimeSourceRunningTime GstRtspclientSinkNtpTimeSource = "running-time" // running-time (2) – Running time based on pipeline clock
	GstRtspclientSinkNtpTimeSourceClockTime GstRtspclientSinkNtpTimeSource = "clock-time" // clock-time (3) – Pipeline clock time
)

