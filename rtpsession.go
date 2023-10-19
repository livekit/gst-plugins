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

type Rtpsession struct {
	Element *gst.Element
}

func NewRtpsession() (*Rtpsession, error) {
	e, err := gst.NewElement("rtpsession")
	if err != nil {
		return nil, err
	}
	return &Rtpsession{Element: e}, nil
}

func NewRtpsessionWithName(name string) (*Rtpsession, error) {
	e, err := gst.NewElementWithName("rtpsession", name)
	if err != nil {
		return nil, err
	}
	return &Rtpsession{Element: e}, nil
}

// ----- Properties -----

// bandwidth (float64)
//
// The bandwidth of the session in bytes per second (0 for auto-discover)

func (e *Rtpsession) GetBandwidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Rtpsession) SetBandwidth(value float64) error {
	return e.Element.SetProperty("bandwidth", value)
}

// internal-session (GstRtpsession)
//
// The internal RTPSession object

func (e *Rtpsession) GetInternalSession() (interface{}, error) {
	return e.Element.GetProperty("internal-session")
}

// max-dropout-time (uint)
//
// The maximum time (milliseconds) of missing packets tolerated.

func (e *Rtpsession) GetMaxDropoutTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-dropout-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpsession) SetMaxDropoutTime(value uint) error {
	return e.Element.SetProperty("max-dropout-time", value)
}

// max-misorder-time (uint)
//
// The maximum time (milliseconds) of misordered packets tolerated.

func (e *Rtpsession) GetMaxMisorderTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-misorder-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpsession) SetMaxMisorderTime(value uint) error {
	return e.Element.SetProperty("max-misorder-time", value)
}

// ntp-time-source (GstRtpNtpTimeSource)
//
// NTP time source for RTCP packets

func (e *Rtpsession) GetNtpTimeSource() (GstRtpNtpTimeSource, error) {
	var value GstRtpNtpTimeSource
	var ok bool
	v, err := e.Element.GetProperty("ntp-time-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtpNtpTimeSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtpNtpTimeSource")
	}
	return value, nil
}

func (e *Rtpsession) SetNtpTimeSource(value GstRtpNtpTimeSource) error {
	e.Element.SetArg("ntp-time-source", string(value))
	return nil
}

// num-active-sources (uint)
//
// The number of active sources in the session

func (e *Rtpsession) GetNumActiveSources() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-active-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// num-sources (uint)
//
// The number of sources in the session

func (e *Rtpsession) GetNumSources() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// probation (uint)
//
// Consecutive packet sequence numbers to accept the source

func (e *Rtpsession) GetProbation() (uint, error) {
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

func (e *Rtpsession) SetProbation(value uint) error {
	return e.Element.SetProperty("probation", value)
}

// rtcp-fraction (float64)
//
// The RTCP bandwidth of the session in bytes per second (or as a real fraction of the RTP bandwidth if < 1.0)

func (e *Rtpsession) GetRtcpFraction() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rtcp-fraction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Rtpsession) SetRtcpFraction(value float64) error {
	return e.Element.SetProperty("rtcp-fraction", value)
}

// rtcp-min-interval (uint64)
//
// Minimum interval between Regular RTCP packet (in ns)

func (e *Rtpsession) GetRtcpMinInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("rtcp-min-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Rtpsession) SetRtcpMinInterval(value uint64) error {
	return e.Element.SetProperty("rtcp-min-interval", value)
}

// rtcp-rr-bandwidth (int)
//
// The RTCP bandwidth used for receivers in bytes per second (-1 = default)

func (e *Rtpsession) GetRtcpRrBandwidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtcp-rr-bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpsession) SetRtcpRrBandwidth(value int) error {
	return e.Element.SetProperty("rtcp-rr-bandwidth", value)
}

// rtcp-rs-bandwidth (int)
//
// The RTCP bandwidth used for senders in bytes per second (-1 = default)

func (e *Rtpsession) GetRtcpRsBandwidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtcp-rs-bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpsession) SetRtcpRsBandwidth(value int) error {
	return e.Element.SetProperty("rtcp-rs-bandwidth", value)
}

// rtcp-sync-send-time (bool)
//
// Use send time or capture time for RTCP sync (TRUE = send time, FALSE = capture time)

func (e *Rtpsession) GetRtcpSyncSendTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync-send-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpsession) SetRtcpSyncSendTime(value bool) error {
	return e.Element.SetProperty("rtcp-sync-send-time", value)
}

// rtp-profile (GstRTPProfile)
//
// RTP profile to use

func (e *Rtpsession) GetRtpProfile() (interface{}, error) {
	return e.Element.GetProperty("rtp-profile")
}

func (e *Rtpsession) SetRtpProfile(value interface{}) error {
	return e.Element.SetProperty("rtp-profile", value)
}

// sdes (GstStructure)
//
// The SDES items of this session

func (e *Rtpsession) GetSdes() (interface{}, error) {
	return e.Element.GetProperty("sdes")
}

func (e *Rtpsession) SetSdes(value interface{}) error {
	return e.Element.SetProperty("sdes", value)
}

// stats (GstStructure)
//
// Various session statistics. This property returns a GstStructure
// with name application/x-rtp-session-stats with the following fields:

// timeout-inactive-sources (bool)
//
// Whether inactive sources should be timed out

func (e *Rtpsession) GetTimeoutInactiveSources() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("timeout-inactive-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpsession) SetTimeoutInactiveSources(value bool) error {
	return e.Element.SetProperty("timeout-inactive-sources", value)
}

// twcc-stats (GstStructure)
//
// Various statistics derived from TWCC. This property returns a GstStructure
// with name RTPTWCCStats with the following fields:

// update-ntp64-header-ext (bool)
//
// Whether RTP NTP header extension should be updated with actual
// NTP time. If not, use the NTP time from buffer timestamp metadata

func (e *Rtpsession) GetUpdateNtp64HeaderExt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("update-ntp64-header-ext")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpsession) SetUpdateNtp64HeaderExt(value bool) error {
	return e.Element.SetProperty("update-ntp64-header-ext", value)
}

// use-pipeline-clock (bool)
//
// Use the pipeline running-time to set the NTP time in the RTCP SR messages (DEPRECATED: Use ntp-time-source property)

func (e *Rtpsession) GetUsePipelineClock() (bool, error) {
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

func (e *Rtpsession) SetUsePipelineClock(value bool) error {
	return e.Element.SetProperty("use-pipeline-clock", value)
}

// ----- Constants -----

type GstRtpNtpTimeSource string

const (
	GstRtpNtpTimeSourceNtp GstRtpNtpTimeSource = "ntp" // ntp (0) – NTP time based on realtime clock
	GstRtpNtpTimeSourceUnix GstRtpNtpTimeSource = "unix" // unix (1) – UNIX time based on realtime clock
	GstRtpNtpTimeSourceRunningTime GstRtpNtpTimeSource = "running-time" // running-time (2) – Running time based on pipeline clock
	GstRtpNtpTimeSourceClockTime GstRtpNtpTimeSource = "clock-time" // clock-time (3) – Pipeline clock time
)

