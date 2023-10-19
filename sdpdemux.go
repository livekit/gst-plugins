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

type Sdpdemux struct {
	Element *gst.Element
}

func NewSdpdemux() (*Sdpdemux, error) {
	e, err := gst.NewElement("sdpdemux")
	if err != nil {
		return nil, err
	}
	return &Sdpdemux{Element: e}, nil
}

func NewSdpdemuxWithName(name string) (*Sdpdemux, error) {
	e, err := gst.NewElementWithName("sdpdemux", name)
	if err != nil {
		return nil, err
	}
	return &Sdpdemux{Element: e}, nil
}

// ----- Properties -----

// debug (bool)
//
// Dump request and response messages to stdout

func (e *Sdpdemux) GetDebug() (bool, error) {
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

func (e *Sdpdemux) SetDebug(value bool) error {
	return e.Element.SetProperty("debug", value)
}

// latency (uint)
//
// Amount of ms to buffer

func (e *Sdpdemux) GetLatency() (uint, error) {
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

func (e *Sdpdemux) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

// media (string)
//
// Media to use, e.g. audio or video (NULL=allow all).

func (e *Sdpdemux) GetMedia() (string, error) {
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

func (e *Sdpdemux) SetMedia(value string) error {
	return e.Element.SetProperty("media", value)
}

// redirect (bool)
//
// Sends a redirection message instead of using a custom session element

func (e *Sdpdemux) GetRedirect() (bool, error) {
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

func (e *Sdpdemux) SetRedirect(value bool) error {
	return e.Element.SetProperty("redirect", value)
}

// rtcp-mode (GstSdpdemuxRtcpmode)
//
// RTCP mode: enable or disable receiving of Sender Reports and
// sending of Receiver Reports.

func (e *Sdpdemux) GetRtcpMode() (GstSdpdemuxRtcpmode, error) {
	var value GstSdpdemuxRtcpmode
	var ok bool
	v, err := e.Element.GetProperty("rtcp-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSdpdemuxRtcpmode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSdpdemuxRtcpmode")
	}
	return value, nil
}

func (e *Sdpdemux) SetRtcpMode(value GstSdpdemuxRtcpmode) error {
	e.Element.SetArg("rtcp-mode", string(value))
	return nil
}

// timeout (uint64)
//
// Fail transport after UDP timeout microseconds (0 = disabled)

func (e *Sdpdemux) GetTimeout() (uint64, error) {
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

func (e *Sdpdemux) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

// timeout-inactive-rtp-sources (bool)
//
// Whether inactive RTP sources in the underlying RTP session
// should be timed out.

func (e *Sdpdemux) GetTimeoutInactiveRtpSources() (bool, error) {
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

func (e *Sdpdemux) SetTimeoutInactiveRtpSources(value bool) error {
	return e.Element.SetProperty("timeout-inactive-rtp-sources", value)
}

// ----- Constants -----

type GstSdpdemuxRtcpmode string

const (
	GstSdpdemuxRtcpmodeSendReceiveRtcp GstSdpdemuxRtcpmode = "Send + Receive RTCP" // Send + Receive RTCP (3) – sendrecv
	GstSdpdemuxRtcpmodeReceiveRtcpSenderReports GstSdpdemuxRtcpmode = "Receive RTCP sender reports" // Receive RTCP sender reports (1) – recvonly
	GstSdpdemuxRtcpmodeSendRtcpReceiverReports GstSdpdemuxRtcpmode = "Send RTCP receiver reports" // Send RTCP receiver reports (2) – sendonly
	GstSdpdemuxRtcpmodeDisableRtcp GstSdpdemuxRtcpmode = "Disable RTCP" // Disable RTCP (0) – inactivate
)

