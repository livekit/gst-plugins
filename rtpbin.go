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

type Rtpbin struct {
	Element *gst.Element
}

func NewRtpbin() (*Rtpbin, error) {
	e, err := gst.NewElement("rtpbin")
	if err != nil {
		return nil, err
	}
	return &Rtpbin{Element: e}, nil
}

func NewRtpbinWithName(name string) (*Rtpbin, error) {
	e, err := gst.NewElementWithName("rtpbin", name)
	if err != nil {
		return nil, err
	}
	return &Rtpbin{Element: e}, nil
}

// ----- Properties -----

// add-reference-timestamp-meta (bool)
//
// When syncing to a RFC7273 clock or after clock synchronization via RTCP or
// inband NTP-64 header extensions has happened, add GstReferenceTimestampMeta
// to buffers with the original reconstructed reference clock timestamp.

func (e *Rtpbin) GetAddReferenceTimestampMeta() (bool, error) {
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

func (e *Rtpbin) SetAddReferenceTimestampMeta(value bool) error {
	return e.Element.SetProperty("add-reference-timestamp-meta", value)
}

// autoremove (bool)
//
// Automatically remove timed out sources

func (e *Rtpbin) GetAutoremove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("autoremove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpbin) SetAutoremove(value bool) error {
	return e.Element.SetProperty("autoremove", value)
}

// buffer-mode (GstRtpjitterBufferMode)
//
// Control the buffering and timestamping mode used by the jitterbuffer.

func (e *Rtpbin) GetBufferMode() (interface{}, error) {
	return e.Element.GetProperty("buffer-mode")
}

func (e *Rtpbin) SetBufferMode(value interface{}) error {
	return e.Element.SetProperty("buffer-mode", value)
}

// do-lost (bool)
//
// Send an event downstream when a packet is lost

func (e *Rtpbin) GetDoLost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-lost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpbin) SetDoLost(value bool) error {
	return e.Element.SetProperty("do-lost", value)
}

// do-retransmission (bool)
//
// Enables RTP retransmission on all streams. To control retransmission on
// a per-SSRC basis, connect to the new-jitterbuffer signal and
// set the do-retransmission property on the
// rtpjitterbuffer object instead.

func (e *Rtpbin) GetDoRetransmission() (bool, error) {
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

func (e *Rtpbin) SetDoRetransmission(value bool) error {
	return e.Element.SetProperty("do-retransmission", value)
}

// do-sync-event (bool)
//
// Send event downstream when a stream is synchronized to the sender

func (e *Rtpbin) GetDoSyncEvent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-sync-event")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpbin) SetDoSyncEvent(value bool) error {
	return e.Element.SetProperty("do-sync-event", value)
}

// drop-on-latency (bool)
//
// Tells the jitterbuffer to never exceed the given latency in size

func (e *Rtpbin) GetDropOnLatency() (bool, error) {
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

func (e *Rtpbin) SetDropOnLatency(value bool) error {
	return e.Element.SetProperty("drop-on-latency", value)
}

// fec-decoders (GstStructure)
//
// Used to provide a factory used to build the FEC decoder for a
// given session, as a command line alternative to
// request-fec-decoder.

// fec-encoders (GstStructure)
//
// Used to provide a factory used to build the FEC encoder for a
// given session, as a command line alternative to
// request-fec-encoder.

// ignore-pt (bool)
//
// Do not demultiplex based on PT values

func (e *Rtpbin) GetIgnorePt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpbin) SetIgnorePt(value bool) error {
	return e.Element.SetProperty("ignore-pt", value)
}

// latency (uint)
//
// Default amount of ms to buffer in the jitterbuffers

func (e *Rtpbin) GetLatency() (uint, error) {
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

func (e *Rtpbin) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

// max-dropout-time (uint)
//
// The maximum time (milliseconds) of missing packets tolerated.

func (e *Rtpbin) GetMaxDropoutTime() (uint, error) {
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

func (e *Rtpbin) SetMaxDropoutTime(value uint) error {
	return e.Element.SetProperty("max-dropout-time", value)
}

// max-misorder-time (uint)
//
// The maximum time (milliseconds) of misordered packets tolerated.

func (e *Rtpbin) GetMaxMisorderTime() (uint, error) {
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

func (e *Rtpbin) SetMaxMisorderTime(value uint) error {
	return e.Element.SetProperty("max-misorder-time", value)
}

// max-rtcp-rtp-time-diff (int)
//
// Maximum amount of time in ms that the RTP time in RTCP SRs is allowed to be ahead (-1 disabled)

func (e *Rtpbin) GetMaxRtcpRtpTimeDiff() (int, error) {
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

func (e *Rtpbin) SetMaxRtcpRtpTimeDiff(value int) error {
	return e.Element.SetProperty("max-rtcp-rtp-time-diff", value)
}

// max-streams (uint)
//
// The maximum number of streams to create for one session

func (e *Rtpbin) GetMaxStreams() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpbin) SetMaxStreams(value uint) error {
	return e.Element.SetProperty("max-streams", value)
}

// max-ts-offset (int64)
//
// Used to set an upper limit of how large a time offset may be. This
// is used to protect against unrealistic values as a result of either
// client,server or clock issues.

func (e *Rtpbin) GetMaxTsOffset() (int64, error) {
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

func (e *Rtpbin) SetMaxTsOffset(value int64) error {
	return e.Element.SetProperty("max-ts-offset", value)
}

// max-ts-offset-adjustment (uint64)
//
// Syncing time stamps to NTP time adds a time offset. This parameter
// specifies the maximum number of nanoseconds per frame that this time offset
// may be adjusted with. This is used to avoid sudden large changes to time
// stamps.

func (e *Rtpbin) GetMaxTsOffsetAdjustment() (uint64, error) {
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

func (e *Rtpbin) SetMaxTsOffsetAdjustment(value uint64) error {
	return e.Element.SetProperty("max-ts-offset-adjustment", value)
}

// min-ts-offset (uint64)
//
// Used to set an lower limit for when a time offset is deemed large enough
// to be useful for sync corrections.

// ntp-sync (bool)
//
// Set the NTP time from the sender reports as the running-time on the
// buffers. When both the sender and receiver have sychronized
// running-time, i.e. when the clock and base-time is shared
// between the receivers and the and the senders, this option can be
// used to synchronize receivers on multiple machines.

func (e *Rtpbin) GetNtpSync() (bool, error) {
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

func (e *Rtpbin) SetNtpSync(value bool) error {
	return e.Element.SetProperty("ntp-sync", value)
}

// ntp-time-source (GstRtpNtpTimeSource)
//
// NTP time source for RTCP packets

func (e *Rtpbin) GetNtpTimeSource() (interface{}, error) {
	return e.Element.GetProperty("ntp-time-source")
}

func (e *Rtpbin) SetNtpTimeSource(value interface{}) error {
	return e.Element.SetProperty("ntp-time-source", value)
}

// rfc7273-sync (bool)
//
// Synchronize received streams to the RFC7273 clock (requires clock and offset to be provided)

func (e *Rtpbin) GetRfc7273Sync() (bool, error) {
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

func (e *Rtpbin) SetRfc7273Sync(value bool) error {
	return e.Element.SetProperty("rfc7273-sync", value)
}

// rtcp-sync (GstRtcpsync)
//
// If not synchronizing (directly) to the NTP clock, determines how to sync
// the various streams.

func (e *Rtpbin) GetRtcpSync() (GstRtcpsync, error) {
	var value GstRtcpsync
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtcpsync)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtcpsync")
	}
	return value, nil
}

func (e *Rtpbin) SetRtcpSync(value GstRtcpsync) error {
	e.Element.SetArg("rtcp-sync", string(value))
	return nil
}

// rtcp-sync-interval (uint)
//
// Determines how often to sync streams using RTCP data or inband NTP-64
// header extensions.

func (e *Rtpbin) GetRtcpSyncInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpbin) SetRtcpSyncInterval(value uint) error {
	return e.Element.SetProperty("rtcp-sync-interval", value)
}

// rtcp-sync-send-time (bool)
//
// Use send time or capture time for RTCP sync (TRUE = send time, FALSE = capture time)

func (e *Rtpbin) GetRtcpSyncSendTime() (bool, error) {
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

func (e *Rtpbin) SetRtcpSyncSendTime(value bool) error {
	return e.Element.SetProperty("rtcp-sync-send-time", value)
}

// rtp-profile (GstRTPProfile)
//
// Sets the default RTP profile of newly created RTP sessions. The
// profile can be changed afterwards on a per-session basis.

func (e *Rtpbin) GetRtpProfile() (interface{}, error) {
	return e.Element.GetProperty("rtp-profile")
}

func (e *Rtpbin) SetRtpProfile(value interface{}) error {
	return e.Element.SetProperty("rtp-profile", value)
}

// sdes (GstStructure)
//
// The SDES items of this session

func (e *Rtpbin) GetSdes() (interface{}, error) {
	return e.Element.GetProperty("sdes")
}

func (e *Rtpbin) SetSdes(value interface{}) error {
	return e.Element.SetProperty("sdes", value)
}

// timeout-inactive-sources (bool)
//
// Whether inactive sources should be timed out

func (e *Rtpbin) GetTimeoutInactiveSources() (bool, error) {
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

func (e *Rtpbin) SetTimeoutInactiveSources(value bool) error {
	return e.Element.SetProperty("timeout-inactive-sources", value)
}

// ts-offset-smoothing-factor (uint)
//
// Controls the weighting between previous and current timestamp offsets in
// a running moving average (RMA):
// ts_offset_average(n) =
// ((ts-offset-smoothing-factor - 1) * ts_offset_average(n - 1) + ts_offset(n)) /
// ts-offset-smoothing-factor

// update-ntp64-header-ext (bool)
//
// Whether RTP NTP header extension should be updated with actual
// NTP time. If not, use the NTP time from buffer timestamp metadata

func (e *Rtpbin) GetUpdateNtp64HeaderExt() (bool, error) {
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

func (e *Rtpbin) SetUpdateNtp64HeaderExt(value bool) error {
	return e.Element.SetProperty("update-ntp64-header-ext", value)
}

// use-pipeline-clock (bool)
//
// Use the pipeline running-time to set the NTP time in the RTCP SR messages (DEPRECATED: Use ntp-time-source property)

func (e *Rtpbin) GetUsePipelineClock() (bool, error) {
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

func (e *Rtpbin) SetUsePipelineClock(value bool) error {
	return e.Element.SetProperty("use-pipeline-clock", value)
}

// ----- Constants -----

type GstRtcpsync string

const (
	GstRtcpsyncAlways GstRtcpsync = "always" // always (0) – always
	GstRtcpsyncInitial GstRtcpsync = "initial" // initial (1) – initial
	GstRtcpsyncRtpInfo GstRtcpsync = "rtp-info" // rtp-info (2) – rtp-info
)

