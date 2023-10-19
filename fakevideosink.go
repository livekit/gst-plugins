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

type Fakevideosink struct {
	Element *gst.Element
}

func NewFakevideosink() (*Fakevideosink, error) {
	e, err := gst.NewElement("fakevideosink")
	if err != nil {
		return nil, err
	}
	return &Fakevideosink{Element: e}, nil
}

func NewFakevideosinkWithName(name string) (*Fakevideosink, error) {
	e, err := gst.NewElementWithName("fakevideosink", name)
	if err != nil {
		return nil, err
	}
	return &Fakevideosink{Element: e}, nil
}

// ----- Properties -----

// allocation-meta-flags (GstFakeVideoSinkAllocationMetaFlags)
//
// Control the behaviour of the sink allocation query handler.

func (e *Fakevideosink) GetAllocationMetaFlags() (GstFakeVideoSinkAllocationMetaFlags, error) {
	var value GstFakeVideoSinkAllocationMetaFlags
	var ok bool
	v, err := e.Element.GetProperty("allocation-meta-flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeVideoSinkAllocationMetaFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeVideoSinkAllocationMetaFlags")
	}
	return value, nil
}

func (e *Fakevideosink) SetAllocationMetaFlags(value GstFakeVideoSinkAllocationMetaFlags) error {
	e.Element.SetArg("allocation-meta-flags", string(value))
	return nil
}

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Fakevideosink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Fakevideosink) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Fakevideosink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// can-activate-pull (bool)
//
// Can activate in pull mode

func (e *Fakevideosink) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// can-activate-push (bool)
//
// Can activate in push mode

func (e *Fakevideosink) GetCanActivatePush() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-push")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

// drop-out-of-segment (bool)
//
// Drop and don't render / hand off out-of-segment buffers

func (e *Fakevideosink) GetDropOutOfSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-out-of-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

// dump (bool)
//
// Dump buffer contents to stdout

func (e *Fakevideosink) GetDump() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dump")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Fakevideosink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// last-message (string)
//
// The message describing current status

func (e *Fakevideosink) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// last-sample (GstSample)
//
// The last sample received in the sink

func (e *Fakevideosink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// max-bitrate (uint64)
//
// The maximum bits per second to render (0 = disabled)

func (e *Fakevideosink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fakevideosink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Fakevideosink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Fakevideosink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// num-buffers (int)
//
// Number of buffers to accept going EOS

func (e *Fakevideosink) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fakevideosink) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Fakevideosink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fakevideosink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Fakevideosink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Fakevideosink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fakevideosink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// signal-handoffs (bool)
//
// Send a signal before unreffing the buffer

func (e *Fakevideosink) GetSignalHandoffs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal-handoffs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

// silent (bool)
//
// Don't produce last_message events

func (e *Fakevideosink) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// state-error (GstFakeVideoSinkStateError)
//
// Generate a state change error

func (e *Fakevideosink) GetStateError() (GstFakeVideoSinkStateError, error) {
	var value GstFakeVideoSinkStateError
	var ok bool
	v, err := e.Element.GetProperty("state-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeVideoSinkStateError)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeVideoSinkStateError")
	}
	return value, nil
}

func (e *Fakevideosink) SetStateError(value GstFakeVideoSinkStateError) error {
	e.Element.SetArg("state-error", string(value))
	return nil
}

// stats (GstStructure)
//
// Sink Statistics

func (e *Fakevideosink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// sync (bool)
//
// Sync on the clock

func (e *Fakevideosink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakevideosink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Fakevideosink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Fakevideosink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Fakevideosink) GetTsOffset() (int64, error) {
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

func (e *Fakevideosink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

// ----- Constants -----

type GstFakeVideoSinkAllocationMetaFlags string

const (
	GstFakeVideoSinkAllocationMetaFlagsCrop GstFakeVideoSinkAllocationMetaFlags = "crop" // crop (0x00000001) – Expose the crop meta as supported
	GstFakeVideoSinkAllocationMetaFlagsOverlayComposition GstFakeVideoSinkAllocationMetaFlags = "overlay-composition" // overlay-composition (0x00000002) – Expose the overlay composition meta as supported
)

type GstFakeVideoSinkStateError string

const (
	GstFakeVideoSinkStateErrorNone GstFakeVideoSinkStateError = "none" // none (0) – No state change errors
	GstFakeVideoSinkStateErrorNullToReady GstFakeVideoSinkStateError = "null-to-ready" // null-to-ready (1) – Fail state change from NULL to READY
	GstFakeVideoSinkStateErrorReadyToPaused GstFakeVideoSinkStateError = "ready-to-paused" // ready-to-paused (2) – Fail state change from READY to PAUSED
	GstFakeVideoSinkStateErrorPausedToPlaying GstFakeVideoSinkStateError = "paused-to-playing" // paused-to-playing (3) – Fail state change from PAUSED to PLAYING
	GstFakeVideoSinkStateErrorPlayingToPaused GstFakeVideoSinkStateError = "playing-to-paused" // playing-to-paused (4) – Fail state change from PLAYING to PAUSED
	GstFakeVideoSinkStateErrorPausedToReady GstFakeVideoSinkStateError = "paused-to-ready" // paused-to-ready (5) – Fail state change from PAUSED to READY
	GstFakeVideoSinkStateErrorReadyToNull GstFakeVideoSinkStateError = "ready-to-null" // ready-to-null (6) – Fail state change from READY to NULL
)

