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

	"github.com/livekit/gstplugins/base"
)

type Fakesink struct {
	*base.GstBaseSink
}

func NewFakesink() (*Fakesink, error) {
	e, err := gst.NewElement("fakesink")
	if err != nil {
		return nil, err
	}
	return &Fakesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewFakesinkWithName(name string) (*Fakesink, error) {
	e, err := gst.NewElementWithName("fakesink", name)
	if err != nil {
		return nil, err
	}
	return &Fakesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// can-activate-pull (bool)
//
// Can activate in pull mode

func (e *Fakesink) GetCanActivatePull() (bool, error) {
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

func (e *Fakesink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// can-activate-push (bool)
//
// Can activate in push mode

func (e *Fakesink) GetCanActivatePush() (bool, error) {
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

func (e *Fakesink) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

// drop-out-of-segment (bool)
//
// Drop and don't render / hand off out-of-segment buffers

func (e *Fakesink) GetDropOutOfSegment() (bool, error) {
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

func (e *Fakesink) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

// dump (bool)
//
// Dump buffer contents to stdout

func (e *Fakesink) GetDump() (bool, error) {
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

func (e *Fakesink) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

// last-message (string)
//
// The message describing current status

func (e *Fakesink) GetLastMessage() (string, error) {
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

// num-buffers (int)
//
// Number of buffers to accept going EOS

func (e *Fakesink) GetNumBuffers() (int, error) {
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

func (e *Fakesink) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// signal-handoffs (bool)
//
// Send a signal before unreffing the buffer

func (e *Fakesink) GetSignalHandoffs() (bool, error) {
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

func (e *Fakesink) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

// silent (bool)
//
// Don't produce last_message events

func (e *Fakesink) GetSilent() (bool, error) {
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

func (e *Fakesink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// state-error (GstFakeSinkStateError)
//
// Generate a state change error

func (e *Fakesink) GetStateError() (GstFakeSinkStateError, error) {
	var value GstFakeSinkStateError
	var ok bool
	v, err := e.Element.GetProperty("state-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSinkStateError)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSinkStateError")
	}
	return value, nil
}

func (e *Fakesink) SetStateError(value GstFakeSinkStateError) error {
	e.Element.SetArg("state-error", string(value))
	return nil
}

// ----- Constants -----

type GstFakeSinkStateError string

const (
	GstFakeSinkStateErrorNone GstFakeSinkStateError = "none" // none (0) – No state change errors
	GstFakeSinkStateErrorNullToReady GstFakeSinkStateError = "null-to-ready" // null-to-ready (1) – Fail state change from NULL to READY
	GstFakeSinkStateErrorReadyToPaused GstFakeSinkStateError = "ready-to-paused" // ready-to-paused (2) – Fail state change from READY to PAUSED
	GstFakeSinkStateErrorPausedToPlaying GstFakeSinkStateError = "paused-to-playing" // paused-to-playing (3) – Fail state change from PAUSED to PLAYING
	GstFakeSinkStateErrorPlayingToPaused GstFakeSinkStateError = "playing-to-paused" // playing-to-paused (4) – Fail state change from PLAYING to PAUSED
	GstFakeSinkStateErrorPausedToReady GstFakeSinkStateError = "paused-to-ready" // paused-to-ready (5) – Fail state change from PAUSED to READY
	GstFakeSinkStateErrorReadyToNull GstFakeSinkStateError = "ready-to-null" // ready-to-null (6) – Fail state change from READY to NULL
)

