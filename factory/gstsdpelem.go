// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Receive data over the network via SDP
type GstSDPDemux struct {
	*gst.Bin
}

func NewSDPDemux() (*GstSDPDemux, error) {
	e, err := gst.NewElement("sdpdemux")
	if err != nil {
		return nil, err
	}
	return &GstSDPDemux{Bin: &gst.Bin{Element: e}}, nil
}

func NewSDPDemuxWithName(name string) (*GstSDPDemux, error) {
	e, err := gst.NewElementWithName("sdpdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstSDPDemux{Bin: &gst.Bin{Element: e}}, nil
}

// Dump request and response messages to stdout
// Default: false
func (e *GstSDPDemux) SetDebug(value bool) error {
	return e.Element.SetProperty("debug", value)
}

func (e *GstSDPDemux) GetDebug() (bool, error) {
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

// Amount of ms to buffer
// Default: 200, Min: 0, Max: -1
func (e *GstSDPDemux) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstSDPDemux) GetLatency() (uint, error) {
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

// Media to use, e.g. audio or video (NULL = all)
// Default: NULL
func (e *GstSDPDemux) SetMedia(value string) error {
	return e.Element.SetProperty("media", value)
}

func (e *GstSDPDemux) GetMedia() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("media")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Sends a redirection message instead of using a custom session element
// Default: true
func (e *GstSDPDemux) SetRedirect(value bool) error {
	return e.Element.SetProperty("redirect", value)
}

func (e *GstSDPDemux) GetRedirect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable or disable receiving of RTCP sender reports and sending of RTCP receiver reports
// Default: Send + Receive RTCP (3)
func (e *GstSDPDemux) SetRtcpMode(value GstSDPDemuxRTCPMode) error {
	e.Element.SetArg("rtcp-mode", string(value))
	return nil
}

func (e *GstSDPDemux) GetRtcpMode() (GstSDPDemuxRTCPMode, error) {
	var value GstSDPDemuxRTCPMode
	var ok bool
	v, err := e.Element.GetProperty("rtcp-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSDPDemuxRTCPMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSDPDemuxRTCPMode")
	}
	return value, nil
}

// Fail transport after UDP timeout microseconds (0 = disabled)
// Default: 10000000, Min: 0, Max: 18446744073709551615
func (e *GstSDPDemux) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstSDPDemux) GetTimeout() (uint64, error) {
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

// Whether RTP sources that don't receive RTP or RTCP packets for longer than 5x RTCP interval should be removed
// Default: true
func (e *GstSDPDemux) SetTimeoutInactiveRtpSources(value bool) error {
	return e.Element.SetProperty("timeout-inactive-rtp-sources", value)
}

func (e *GstSDPDemux) GetTimeoutInactiveRtpSources() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("timeout-inactive-rtp-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Stream RTP based on an SDP
type GstSdpSrc struct {
	*gst.Bin
}

func NewSdpSrc() (*GstSdpSrc, error) {
	e, err := gst.NewElement("sdpsrc")
	if err != nil {
		return nil, err
	}
	return &GstSdpSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewSdpSrcWithName(name string) (*GstSdpSrc, error) {
	e, err := gst.NewElementWithName("sdpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSdpSrc{Bin: &gst.Bin{Element: e}}, nil
}

// SDP description used instead of location
// Default: NULL
func (e *GstSdpSrc) SetSdp(value string) error {
	return e.Element.SetProperty("sdp", value)
}

func (e *GstSdpSrc) GetSdp() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sdp")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// URI to SDP file (sdp:///path/to/file)
// Default: NULL
func (e *GstSdpSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstSdpSrc) GetLocation() (string, error) {
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

type GstSDPDemuxRTCPMode string

const (
	GstSDPDemuxRTCPModeSendReceiveRTCP          GstSDPDemuxRTCPMode = "Send + Receive RTCP"         // sendrecv
	GstSDPDemuxRTCPModeReceiveRTCPSenderReports GstSDPDemuxRTCPMode = "Receive RTCP sender reports" // recvonly
	GstSDPDemuxRTCPModeSendRTCPReceiverReports  GstSDPDemuxRTCPMode = "Send RTCP receiver reports"  // sendonly
	GstSDPDemuxRTCPModeDisableRTCP              GstSDPDemuxRTCPMode = "Disable RTCP"                // inactivate
)
