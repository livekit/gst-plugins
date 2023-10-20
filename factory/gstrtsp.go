// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Receive data over the network via RTSP (RFC 2326)
type GstRTSPSrc struct {
	*gst.Bin
}

func NewRTSPSrc() (*GstRTSPSrc, error) {
	e, err := gst.NewElement("rtspsrc")
	if err != nil {
		return nil, err
	}
	return &GstRTSPSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewRTSPSrcWithName(name string) (*GstRTSPSrc, error) {
	e, err := gst.NewElementWithName("rtspsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRTSPSrc{Bin: &gst.Bin{Element: e}}, nil
}

// Whether to ignore the x-server-ip-address server header reply
// Default: false
func (e *GstRTSPSrc) SetIgnoreXServerReply(value bool) error {
	return e.Element.SetProperty("ignore-x-server-reply", value)
}

func (e *GstRTSPSrc) GetIgnoreXServerReply() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-x-server-reply")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The network interface on which to join the multicast group
// Default: NULL
func (e *GstRTSPSrc) SetMulticastIface(value string) error {
	return e.Element.SetProperty("multicast-iface", value)
}

func (e *GstRTSPSrc) GetMulticastIface() (string, error) {
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

// NTP time source for RTCP packets
// Default: ntp (0)
func (e *GstRTSPSrc) SetNtpTimeSource(value GstRTSPSrcNtpTimeSource) error {
	e.Element.SetArg("ntp-time-source", string(value))
	return nil
}

func (e *GstRTSPSrc) GetNtpTimeSource() (GstRTSPSrcNtpTimeSource, error) {
	var value GstRTSPSrcNtpTimeSource
	var ok bool
	v, err := e.Element.GetProperty("ntp-time-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRTSPSrcNtpTimeSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRTSPSrcNtpTimeSource")
	}
	return value, nil
}

// Act as an ONVIF client
// Default: false
func (e *GstRTSPSrc) SetOnvifMode(value bool) error {
	return e.Element.SetProperty("onvif-mode", value)
}

func (e *GstRTSPSrc) GetOnvifMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("onvif-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// HTTP proxy URI user password for authentication
// Default: NULL
func (e *GstRTSPSrc) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

func (e *GstRTSPSrc) GetProxyPw() (string, error) {
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

// Dump request and response messages to stdout(DEPRECATED: Printed all RTSP message to gstreamer log as 'log' level)
// Default: false
func (e *GstRTSPSrc) SetDebug(value bool) error {
	return e.Element.SetProperty("debug", value)
}

func (e *GstRTSPSrc) GetDebug() (bool, error) {
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

// The RTSP version that should be tried first when negotiating version.
// Default: 1-0 (16)
func (e *GstRTSPSrc) SetDefaultRtspVersion(value interface{}) error {
	return e.Element.SetProperty("default-rtsp-version", value)
}

func (e *GstRTSPSrc) GetDefaultRtspVersion() (interface{}, error) {
	return e.Element.GetProperty("default-rtsp-version")
}

// Send RTSP keep alive packets, disable for old incompatible server.
// Default: true
func (e *GstRTSPSrc) SetDoRtspKeepAlive(value bool) error {
	return e.Element.SetProperty("do-rtsp-keep-alive", value)
}

func (e *GstRTSPSrc) GetDoRtspKeepAlive() (bool, error) {
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

// Max number of retries when allocating RTP ports.
// Default: 20, Min: 0, Max: 65535
func (e *GstRTSPSrc) SetRetry(value uint) error {
	return e.Element.SetProperty("retry", value)
}

func (e *GstRTSPSrc) GetRetry() (uint, error) {
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

// TLS database with anchor certificate authorities used to validate the server certificate

func (e *GstRTSPSrc) SetTlsDatabase(value interface{}) error {
	return e.Element.SetProperty("tls-database", value)
}

func (e *GstRTSPSrc) GetTlsDatabase() (interface{}, error) {
	return e.Element.GetProperty("tls-database")
}

// Size of the kernel UDP receive buffer in bytes, 0=default
// Default: 524288, Min: 0, Max: 2147483647
func (e *GstRTSPSrc) SetUdpBufferSize(value int) error {
	return e.Element.SetProperty("udp-buffer-size", value)
}

func (e *GstRTSPSrc) GetUdpBufferSize() (int, error) {
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

// Send RTCP packets, disable for old incompatible server.
// Default: true
func (e *GstRTSPSrc) SetDoRtcp(value bool) error {
	return e.Element.SetProperty("do-rtcp", value)
}

func (e *GstRTSPSrc) GetDoRtcp() (bool, error) {
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

// The maximum absolute value of the time offset in (nanoseconds). Note, if the ntp-sync parameter is set the default value is changed to 0 (no limit)
// Default: 3000000000, Min: 0, Max: 9223372036854775807
func (e *GstRTSPSrc) SetMaxTsOffset(value int64) error {
	return e.Element.SetProperty("max-ts-offset", value)
}

func (e *GstRTSPSrc) GetMaxTsOffset() (int64, error) {
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

// A GTlsInteraction object to prompt the user for password or certificate

func (e *GstRTSPSrc) SetTlsInteraction(value interface{}) error {
	return e.Element.SetProperty("tls-interaction", value)
}

func (e *GstRTSPSrc) GetTlsInteraction() (interface{}, error) {
	return e.Element.GetProperty("tls-interaction")
}

// The maximum number of nanoseconds per frame that time stamp offsets may be adjusted (0 = no limit).
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstRTSPSrc) SetMaxTsOffsetAdjustment(value uint64) error {
	return e.Element.SetProperty("max-ts-offset-adjustment", value)
}

func (e *GstRTSPSrc) GetMaxTsOffsetAdjustment() (uint64, error) {
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

// Allowed lower transport protocols
// Default: tcp+udp-mcast+udp
func (e *GstRTSPSrc) SetProtocols(value interface{}) error {
	return e.Element.SetProperty("protocols", value)
}

func (e *GstRTSPSrc) GetProtocols() (interface{}, error) {
	return e.Element.GetProperty("protocols")
}

// HTTP proxy URI user id for authentication
// Default: NULL
func (e *GstRTSPSrc) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

func (e *GstRTSPSrc) GetProxyId() (string, error) {
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

// Only send the basic RTSP headers for broken encoders
// Default: false
func (e *GstRTSPSrc) SetShortHeader(value bool) error {
	return e.Element.SetProperty("short-header", value)
}

func (e *GstRTSPSrc) GetShortHeader() (bool, error) {
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

// Use the pipeline running-time to set the NTP time in the RTCP SR messages(DEPRECATED: Use ntp-time-source property)
// Default: false
func (e *GstRTSPSrc) SetUsePipelineClock(value bool) error {
	return e.Element.SetProperty("use-pipeline-clock", value)
}

func (e *GstRTSPSrc) GetUsePipelineClock() (bool, error) {
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

// Add Reference Timestamp Meta to buffers with the original clock timestamp before any adjustments when syncing to an RFC7273 clock.
// Default: false
func (e *GstRTSPSrc) SetAddReferenceTimestampMeta(value bool) error {
	return e.Element.SetProperty("add-reference-timestamp-meta", value)
}

func (e *GstRTSPSrc) GetAddReferenceTimestampMeta() (bool, error) {
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

// Amount of ms to buffer
// Default: 2000, Min: 0, Max: -1
func (e *GstRTSPSrc) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstRTSPSrc) GetLatency() (uint, error) {
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

// Method to use for traversing firewalls and NAT
// Default: dummy (1)
func (e *GstRTSPSrc) SetNatMethod(value GstRTSPNatMethod) error {
	e.Element.SetArg("nat-method", string(value))
	return nil
}

func (e *GstRTSPSrc) GetNatMethod() (GstRTSPNatMethod, error) {
	var value GstRTSPNatMethod
	var ok bool
	v, err := e.Element.GetProperty("nat-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRTSPNatMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRTSPNatMethod")
	}
	return value, nil
}

// Consecutive packet sequence numbers to accept the source
// Default: 2, Min: 0, Max: -1
func (e *GstRTSPSrc) SetProbation(value uint) error {
	return e.Element.SetProperty("probation", value)
}

func (e *GstRTSPSrc) GetProbation() (uint, error) {
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

// Retry TCP transport after UDP timeout microseconds (0 = disabled)
// Default: 5000000, Min: 0, Max: 18446744073709551615
func (e *GstRTSPSrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstRTSPSrc) GetTimeout() (uint64, error) {
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

// Reconnect to the server if RTSP connection is closed when doing UDP
// Default: true
func (e *GstRTSPSrc) SetUdpReconnect(value bool) error {
	return e.Element.SetProperty("udp-reconnect", value)
}

func (e *GstRTSPSrc) GetUdpReconnect() (bool, error) {
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

// Tells the jitterbuffer to never exceed the given latency in size
// Default: false
func (e *GstRTSPSrc) SetDropOnLatency(value bool) error {
	return e.Element.SetProperty("drop-on-latency", value)
}

func (e *GstRTSPSrc) GetDropOnLatency() (bool, error) {
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

// Whether to act as a live source
// Default: true
func (e *GstRTSPSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstRTSPSrc) GetIsLive() (bool, error) {
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

// Maximum amount of time in ms that the RTP time in RTCP SRs is allowed to be ahead (-1 disabled)
// Default: 1000, Min: -1, Max: 2147483647
func (e *GstRTSPSrc) SetMaxRtcpRtpTimeDiff(value int) error {
	return e.Element.SetProperty("max-rtcp-rtp-time-diff", value)
}

func (e *GstRTSPSrc) GetMaxRtcpRtpTimeDiff() (int, error) {
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

// Client port range that can be used to receive RTP and RTCP data, eg. 3000-3005 (NULL = no restrictions)
// Default: NULL
func (e *GstRTSPSrc) SetPortRange(value string) error {
	return e.Element.SetProperty("port-range", value)
}

func (e *GstRTSPSrc) GetPortRange() (string, error) {
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

// When transitioning PAUSED-READY, allow up to timeout (in nanoseconds) delay in order to send teardown (0 = disabled)
// Default: 100000000, Min: 0, Max: 18446744073709551615
func (e *GstRTSPSrc) SetTeardownTimeout(value uint64) error {
	return e.Element.SetProperty("teardown-timeout", value)
}

func (e *GstRTSPSrc) GetTeardownTimeout() (uint64, error) {
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

// Location of the RTSP url to read
// Default: NULL
func (e *GstRTSPSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstRTSPSrc) GetLocation() (string, error) {
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

// Synchronize received streams to the NTP clock
// Default: false
func (e *GstRTSPSrc) SetNtpSync(value bool) error {
	return e.Element.SetProperty("ntp-sync", value)
}

func (e *GstRTSPSrc) GetNtpSync() (bool, error) {
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

// When in onvif-mode, whether to set Rate-Control to yes or no
// Default: true
func (e *GstRTSPSrc) SetOnvifRateControl(value bool) error {
	return e.Element.SetProperty("onvif-rate-control", value)
}

func (e *GstRTSPSrc) GetOnvifRateControl() (bool, error) {
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

// Proxy settings for HTTP tunneling. Format: [http://][user:passwd@]host[:port]
// Default: NULL
func (e *GstRTSPSrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

func (e *GstRTSPSrc) GetProxy() (string, error) {
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

// Fail after timeout microseconds on TCP connections (0 = disabled)
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstRTSPSrc) SetTcpTimeout(value uint64) error {
	return e.Element.SetProperty("tcp-timeout", value)
}

func (e *GstRTSPSrc) GetTcpTimeout() (uint64, error) {
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

// The User-Agent string to send to the server
// Default: GStreamer/{VERSION}
func (e *GstRTSPSrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

func (e *GstRTSPSrc) GetUserAgent() (string, error) {
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

// RTSP location URI user id for authentication
// Default: NULL
func (e *GstRTSPSrc) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

func (e *GstRTSPSrc) GetUserId() (string, error) {
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

// The type of backchannel to setup. Default is 'none'.
// Default: none (0)
func (e *GstRTSPSrc) SetBackchannel(value GstRTSPBackchannel) error {
	e.Element.SetArg("backchannel", string(value))
	return nil
}

func (e *GstRTSPSrc) GetBackchannel() (GstRTSPBackchannel, error) {
	var value GstRTSPBackchannel
	var ok bool
	v, err := e.Element.GetProperty("backchannel")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRTSPBackchannel)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRTSPBackchannel")
	}
	return value, nil
}

// Control the buffering algorithm in use
// Default: auto (3)
func (e *GstRTSPSrc) SetBufferMode(value GstRTSPSrcBufferMode) error {
	e.Element.SetArg("buffer-mode", string(value))
	return nil
}

func (e *GstRTSPSrc) GetBufferMode() (GstRTSPSrcBufferMode, error) {
	var value GstRTSPSrcBufferMode
	var ok bool
	v, err := e.Element.GetProperty("buffer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRTSPSrcBufferMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRTSPSrcBufferMode")
	}
	return value, nil
}

// Ask the server to retransmit lost packets
// Default: true
func (e *GstRTSPSrc) SetDoRetransmission(value bool) error {
	return e.Element.SetProperty("do-retransmission", value)
}

func (e *GstRTSPSrc) GetDoRetransmission() (bool, error) {
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

// RTP package size to suggest to server (0 = disabled)
// Default: 0, Min: 0, Max: 65536
func (e *GstRTSPSrc) SetRtpBlocksize(value uint) error {
	return e.Element.SetProperty("rtp-blocksize", value)
}

func (e *GstRTSPSrc) GetRtpBlocksize() (uint, error) {
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

// The SDES items of this session

func (e *GstRTSPSrc) SetSdes(value *gst.Structure) error {
	return e.Element.SetProperty("sdes", value)
}

func (e *GstRTSPSrc) GetSdes() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("sdes")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// TLS certificate validation flags used to validate the server certificate
// Default: validate-all
func (e *GstRTSPSrc) SetTlsValidationFlags(value interface{}) error {
	return e.Element.SetProperty("tls-validation-flags", value)
}

func (e *GstRTSPSrc) GetTlsValidationFlags() (interface{}, error) {
	return e.Element.GetProperty("tls-validation-flags")
}

// RTSP location URI user password for authentication
// Default: NULL
func (e *GstRTSPSrc) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

func (e *GstRTSPSrc) GetUserPw() (string, error) {
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

// Network connection speed in kbps (0 = unknown)
// Default: 0, Min: 0, Max: 18446744073709551
func (e *GstRTSPSrc) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

func (e *GstRTSPSrc) GetConnectionSpeed() (uint64, error) {
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

// Extra headers to append to HTTP requests when in tunneled mode
// Default: extra-http-request-headers;
func (e *GstRTSPSrc) SetExtraHttpRequestHeaders(value *gst.Structure) error {
	return e.Element.SetProperty("extra-http-request-headers", value)
}

func (e *GstRTSPSrc) GetExtraHttpRequestHeaders() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("extra-http-request-headers")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Synchronize received streams to the RFC7273 clock (requires clock and offset to be provided)
// Default: false
func (e *GstRTSPSrc) SetRfc7273Sync(value bool) error {
	return e.Element.SetProperty("rfc7273-sync", value)
}

func (e *GstRTSPSrc) GetRfc7273Sync() (bool, error) {
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

// Accepts raw RTP and RTCP packets and sends them forward
type GstRTPDec struct {
	*gst.Element
}

func NewRTPDec() (*GstRTPDec, error) {
	e, err := gst.NewElement("rtpdec")
	if err != nil {
		return nil, err
	}
	return &GstRTPDec{Element: e}, nil
}

func NewRTPDecWithName(name string) (*GstRTPDec, error) {
	e, err := gst.NewElementWithName("rtpdec", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPDec{Element: e}, nil
}

// Amount of ms to buffer
// Default: 200, Min: 0, Max: -1
func (e *GstRTPDec) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstRTPDec) GetLatency() (uint, error) {
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

type GstRTSPNatMethod string

const (
	GstRTSPNatMethodNone  GstRTSPNatMethod = "none"  // None
	GstRTSPNatMethodDummy GstRTSPNatMethod = "dummy" // Send Dummy packets
)

type GstRTSPSrcBufferMode string

const (
	GstRTSPSrcBufferModeNone   GstRTSPSrcBufferMode = "none"   // Only use RTP timestamps
	GstRTSPSrcBufferModeSlave  GstRTSPSrcBufferMode = "slave"  // Slave receiver to sender clock
	GstRTSPSrcBufferModeBuffer GstRTSPSrcBufferMode = "buffer" // Do low/high watermark buffering
	GstRTSPSrcBufferModeAuto   GstRTSPSrcBufferMode = "auto"   // Choose mode depending on stream live
	GstRTSPSrcBufferModeSynced GstRTSPSrcBufferMode = "synced" // Synchronized sender and receiver clocks
)

type GstRTSPSrcNtpTimeSource string

const (
	GstRTSPSrcNtpTimeSourceNtp         GstRTSPSrcNtpTimeSource = "ntp"          // NTP time based on realtime clock
	GstRTSPSrcNtpTimeSourceUnix        GstRTSPSrcNtpTimeSource = "unix"         // UNIX time based on realtime clock
	GstRTSPSrcNtpTimeSourceRunningTime GstRTSPSrcNtpTimeSource = "running-time" // Running time based on pipeline clock
	GstRTSPSrcNtpTimeSourceClockTime   GstRTSPSrcNtpTimeSource = "clock-time"   // Pipeline clock time
)

type GstRTSPBackchannel string

const (
	GstRTSPBackchannelNone  GstRTSPBackchannel = "none"  // No backchannel
	GstRTSPBackchannelOnvif GstRTSPBackchannel = "onvif" // ONVIF audio backchannel
)
