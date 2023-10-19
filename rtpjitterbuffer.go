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

type Rtpjitterbuffer struct {
	Element *gst.Element
}

func NewRtpjitterbuffer() (*Rtpjitterbuffer, error) {
	e, err := gst.NewElement("rtpjitterbuffer")
	if err != nil {
		return nil, err
	}
	return &Rtpjitterbuffer{Element: e}, nil
}

func NewRtpjitterbufferWithName(name string) (*Rtpjitterbuffer, error) {
	e, err := gst.NewElementWithName("rtpjitterbuffer", name)
	if err != nil {
		return nil, err
	}
	return &Rtpjitterbuffer{Element: e}, nil
}

// ----- Properties -----

// add-reference-timestamp-meta (bool)
//
// When syncing to a RFC7273 clock or after clock synchronization via RTCP or
// inband NTP-64 header extensions has happened, add GstReferenceTimestampMeta
// to buffers with the original reconstructed reference clock timestamp.

func (e *Rtpjitterbuffer) GetAddReferenceTimestampMeta() (bool, error) {
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

func (e *Rtpjitterbuffer) SetAddReferenceTimestampMeta(value bool) error {
	return e.Element.SetProperty("add-reference-timestamp-meta", value)
}

// do-lost (bool)
//
// Send out a GstRTPPacketLost event downstream when a packet is considered
// lost.

func (e *Rtpjitterbuffer) GetDoLost() (bool, error) {
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

func (e *Rtpjitterbuffer) SetDoLost(value bool) error {
	return e.Element.SetProperty("do-lost", value)
}

// do-retransmission (bool)
//
// Send out a GstRTPRetransmission event upstream when a packet is considered
// late and should be retransmitted.

func (e *Rtpjitterbuffer) GetDoRetransmission() (bool, error) {
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

func (e *Rtpjitterbuffer) SetDoRetransmission(value bool) error {
	return e.Element.SetProperty("do-retransmission", value)
}

// drop-messages-interval (uint)
//
// Minimal time in milliseconds between posting dropped packet messages, if enabled
// by setting property by setting post-drop-messages to TRUE.
// If interval is set to 0, every dropped packet will result in a drop message being posted.

func (e *Rtpjitterbuffer) GetDropMessagesInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("drop-messages-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetDropMessagesInterval(value uint) error {
	return e.Element.SetProperty("drop-messages-interval", value)
}

// drop-on-latency (bool)
//
// Drop oldest buffers when the queue is completely filled.

func (e *Rtpjitterbuffer) GetDropOnLatency() (bool, error) {
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

func (e *Rtpjitterbuffer) SetDropOnLatency(value bool) error {
	return e.Element.SetProperty("drop-on-latency", value)
}

// faststart-min-packets (uint)
//
// The number of consecutive packets needed to start (set to 0 to
// disable faststart. The jitterbuffer will by default start after the
// latency has elapsed)

func (e *Rtpjitterbuffer) GetFaststartMinPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("faststart-min-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetFaststartMinPackets(value uint) error {
	return e.Element.SetProperty("faststart-min-packets", value)
}

// latency (uint)
//
// The maximum latency of the jitterbuffer. Packets will be kept in the buffer
// for at most this time.

func (e *Rtpjitterbuffer) GetLatency() (uint, error) {
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

func (e *Rtpjitterbuffer) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

// max-dropout-time (uint)
//
// The maximum time (milliseconds) of missing packets tolerated.

func (e *Rtpjitterbuffer) GetMaxDropoutTime() (uint, error) {
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

func (e *Rtpjitterbuffer) SetMaxDropoutTime(value uint) error {
	return e.Element.SetProperty("max-dropout-time", value)
}

// max-misorder-time (uint)
//
// The maximum time (milliseconds) of misordered packets tolerated.

func (e *Rtpjitterbuffer) GetMaxMisorderTime() (uint, error) {
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

func (e *Rtpjitterbuffer) SetMaxMisorderTime(value uint) error {
	return e.Element.SetProperty("max-misorder-time", value)
}

// max-rtcp-rtp-time-diff (int)
//
// The maximum amount of time in ms that the RTP time in the RTCP SRs
// is allowed to be ahead of the last RTP packet we received. Use
// -1 to disable ignoring of RTCP packets.

func (e *Rtpjitterbuffer) GetMaxRtcpRtpTimeDiff() (int, error) {
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

func (e *Rtpjitterbuffer) SetMaxRtcpRtpTimeDiff(value int) error {
	return e.Element.SetProperty("max-rtcp-rtp-time-diff", value)
}

// max-ts-offset-adjustment (uint64)
//
// The maximum number of nanoseconds per frame that time offset may be
// adjusted with. This is used to avoid sudden large changes to time stamps.

func (e *Rtpjitterbuffer) GetMaxTsOffsetAdjustment() (uint64, error) {
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

func (e *Rtpjitterbuffer) SetMaxTsOffsetAdjustment(value uint64) error {
	return e.Element.SetProperty("max-ts-offset-adjustment", value)
}

// mode (GstRtpjitterBufferMode)
//
// Control the buffering and timestamping mode used by the jitterbuffer.

func (e *Rtpjitterbuffer) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *Rtpjitterbuffer) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// percent (int)
//
// The percent of the jitterbuffer that is filled.

func (e *Rtpjitterbuffer) GetPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// post-drop-messages (bool)
//
// Post custom messages to the bus when a packet is dropped by the
// jitterbuffer due to arriving too late, being already considered lost,
// or being dropped due to the drop-on-latency property being enabled.
// Message is of type GST_MESSAGE_ELEMENT and contains a GstStructure named
// "drop-msg" with the following fields:

// rfc7273-sync (bool)
//
// Synchronize received streams to the RFC7273 clock (requires clock and offset to be provided)

func (e *Rtpjitterbuffer) GetRfc7273Sync() (bool, error) {
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

func (e *Rtpjitterbuffer) SetRfc7273Sync(value bool) error {
	return e.Element.SetProperty("rfc7273-sync", value)
}

// rtx-deadline (int)
//
// The deadline for a valid RTX request in ms.

// rtx-delay (int)
//
// When a packet did not arrive at the expected time, wait this extra amount
// of time before sending a retransmission event.

// rtx-delay-reorder (int)
//
// Assume that a retransmission event should be sent when we see
// this much packet reordering.

// rtx-max-retries (int)
//
// The maximum number of retries to request a retransmission.

// rtx-min-delay (uint)
//
// When a packet did not arrive at the expected time, wait at least this extra amount
// of time before sending a retransmission event.

func (e *Rtpjitterbuffer) GetRtxMinDelay() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtx-min-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetRtxMinDelay(value uint) error {
	return e.Element.SetProperty("rtx-min-delay", value)
}

// rtx-min-retry-timeout (int)
//
// The minimum amount of time between retry timeouts. When
// GstRtpJitterBuffer::rtx-retry-timeout is -1, this value ensures a
// minimum interval between retry timeouts.

// rtx-next-seqnum (bool)
//
// Estimate when the next packet should arrive and schedule a retransmission
// request for it.
// This is, when packet N arrives, a GstRTPRetransmission event is schedule
// for packet N+1. So it will be requested if it does not arrive at the expected time.
// The expected time is calculated using the dts of N and the packet spacing.

func (e *Rtpjitterbuffer) GetRtxNextSeqnum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtx-next-seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetRtxNextSeqnum(value bool) error {
	return e.Element.SetProperty("rtx-next-seqnum", value)
}

// rtx-retry-period (int)
//
// The amount of time to try to get a retransmission.

// rtx-retry-timeout (int)
//
// When no packet has been received after sending a retransmission event
// for this time, retry sending a retransmission event.

// rtx-stats-timeout (uint)
//
// The time to wait for a retransmitted packet after it has been
// considered lost in order to collect RTX statistics.

func (e *Rtpjitterbuffer) GetRtxStatsTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtx-stats-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetRtxStatsTimeout(value uint) error {
	return e.Element.SetProperty("rtx-stats-timeout", value)
}

// stats (GstStructure)
//
// Various jitterbuffer statistics. This property returns a GstStructure
// with name application/x-rtp-jitterbuffer-stats with the following fields:

// sync-interval (uint)
//
// Determines how often to sync streams using RTCP data or inband NTP-64
// header extensions.

func (e *Rtpjitterbuffer) GetSyncInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sync-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetSyncInterval(value uint) error {
	return e.Element.SetProperty("sync-interval", value)
}

// ts-offset (int64)
//
// Adjust GStreamer output buffer timestamps in the jitterbuffer with offset.
// This is mainly used to ensure interstream synchronisation.

func (e *Rtpjitterbuffer) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Rtpjitterbuffer) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

// ----- Constants -----

type RtpjitterBufferMode string

const (
	RtpjitterBufferModeNone RtpjitterBufferMode = "none" // none (0) – Only use RTP timestamps
	RtpjitterBufferModeSlave RtpjitterBufferMode = "slave" // slave (1) – Slave receiver to sender clock
	RtpjitterBufferModeBuffer RtpjitterBufferMode = "buffer" // buffer (2) – Do low/high watermark buffering
	RtpjitterBufferModeSynced RtpjitterBufferMode = "synced" // synced (4) – Synchronized sender and receiver clocks
)

