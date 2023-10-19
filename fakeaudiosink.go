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

type Fakeaudiosink struct {
	Element *gst.Element
}

func NewFakeaudiosink() (*Fakeaudiosink, error) {
	e, err := gst.NewElement("fakeaudiosink")
	if err != nil {
		return nil, err
	}
	return &Fakeaudiosink{Element: e}, nil
}

func NewFakeaudiosinkWithName(name string) (*Fakeaudiosink, error) {
	e, err := gst.NewElementWithName("fakeaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &Fakeaudiosink{Element: e}, nil
}

// ----- Properties -----

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Fakeaudiosink) GetAsync() (bool, error) {
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

func (e *Fakeaudiosink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Fakeaudiosink) GetBlocksize() (uint, error) {
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

func (e *Fakeaudiosink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// can-activate-pull (bool)
//
// Can activate in pull mode

func (e *Fakeaudiosink) GetCanActivatePull() (bool, error) {
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

func (e *Fakeaudiosink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// can-activate-push (bool)
//
// Can activate in push mode

func (e *Fakeaudiosink) GetCanActivatePush() (bool, error) {
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

func (e *Fakeaudiosink) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

// drop-out-of-segment (bool)
//
// Drop and don't render / hand off out-of-segment buffers

func (e *Fakeaudiosink) GetDropOutOfSegment() (bool, error) {
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

func (e *Fakeaudiosink) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

// dump (bool)
//
// Dump buffer contents to stdout

func (e *Fakeaudiosink) GetDump() (bool, error) {
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

func (e *Fakeaudiosink) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Fakeaudiosink) GetEnableLastSample() (bool, error) {
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

func (e *Fakeaudiosink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// last-message (string)
//
// The message describing current status

func (e *Fakeaudiosink) GetLastMessage() (string, error) {
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

func (e *Fakeaudiosink) GetLastSample() (*gst.Sample, error) {
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

func (e *Fakeaudiosink) GetMaxBitrate() (uint64, error) {
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

func (e *Fakeaudiosink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Fakeaudiosink) GetMaxLateness() (int64, error) {
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

func (e *Fakeaudiosink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// mute (bool)
//
// Control the mute state

func (e *Fakeaudiosink) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakeaudiosink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// num-buffers (int)
//
// Number of buffers to accept going EOS

func (e *Fakeaudiosink) GetNumBuffers() (int, error) {
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

func (e *Fakeaudiosink) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Fakeaudiosink) GetProcessingDeadline() (uint64, error) {
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

func (e *Fakeaudiosink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Fakeaudiosink) GetQos() (bool, error) {
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

func (e *Fakeaudiosink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Fakeaudiosink) GetRenderDelay() (uint64, error) {
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

func (e *Fakeaudiosink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// signal-handoffs (bool)
//
// Send a signal before unreffing the buffer

func (e *Fakeaudiosink) GetSignalHandoffs() (bool, error) {
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

func (e *Fakeaudiosink) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

// silent (bool)
//
// Don't produce last_message events

func (e *Fakeaudiosink) GetSilent() (bool, error) {
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

func (e *Fakeaudiosink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// state-error (GstFakeAudioSinkStateError)
//
// Generate a state change error

func (e *Fakeaudiosink) GetStateError() (GstFakeAudioSinkStateError, error) {
	var value GstFakeAudioSinkStateError
	var ok bool
	v, err := e.Element.GetProperty("state-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeAudioSinkStateError)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeAudioSinkStateError")
	}
	return value, nil
}

func (e *Fakeaudiosink) SetStateError(value GstFakeAudioSinkStateError) error {
	e.Element.SetArg("state-error", string(value))
	return nil
}

// stats (GstStructure)
//
// Sink Statistics

func (e *Fakeaudiosink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// sync (bool)
//
// Sync on the clock

func (e *Fakeaudiosink) GetSync() (bool, error) {
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

func (e *Fakeaudiosink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Fakeaudiosink) GetThrottleTime() (uint64, error) {
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

func (e *Fakeaudiosink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Fakeaudiosink) GetTsOffset() (int64, error) {
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

func (e *Fakeaudiosink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

// volume (float64)
//
// Control the audio volume

func (e *Fakeaudiosink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Fakeaudiosink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

// ----- Constants -----

type GstFakeAudioSinkStateError string

const (
	GstFakeAudioSinkStateErrorNone GstFakeAudioSinkStateError = "none" // none (0) – No state change errors
	GstFakeAudioSinkStateErrorNullToReady GstFakeAudioSinkStateError = "null-to-ready" // null-to-ready (1) – Fail state change from NULL to READY
	GstFakeAudioSinkStateErrorReadyToPaused GstFakeAudioSinkStateError = "ready-to-paused" // ready-to-paused (2) – Fail state change from READY to PAUSED
	GstFakeAudioSinkStateErrorPausedToPlaying GstFakeAudioSinkStateError = "paused-to-playing" // paused-to-playing (3) – Fail state change from PAUSED to PLAYING
	GstFakeAudioSinkStateErrorPlayingToPaused GstFakeAudioSinkStateError = "playing-to-paused" // playing-to-paused (4) – Fail state change from PLAYING to PAUSED
	GstFakeAudioSinkStateErrorPausedToReady GstFakeAudioSinkStateError = "paused-to-ready" // paused-to-ready (5) – Fail state change from PAUSED to READY
	GstFakeAudioSinkStateErrorReadyToNull GstFakeAudioSinkStateError = "ready-to-null" // ready-to-null (6) – Fail state change from READY to NULL
)

