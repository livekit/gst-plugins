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

type Identity struct {
	*base.GstBaseTransform
}

func NewIdentity() (*Identity, error) {
	e, err := gst.NewElement("identity")
	if err != nil {
		return nil, err
	}
	return &Identity{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewIdentityWithName(name string) (*Identity, error) {
	e, err := gst.NewElementWithName("identity", name)
	if err != nil {
		return nil, err
	}
	return &Identity{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// check-imperfect-offset (bool)
//
// Send element messages if offset and offset_end do not match up

func (e *Identity) GetCheckImperfectOffset() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("check-imperfect-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Identity) SetCheckImperfectOffset(value bool) error {
	return e.Element.SetProperty("check-imperfect-offset", value)
}

// check-imperfect-timestamp (bool)
//
// Send element messages if timestamps and durations do not match up

func (e *Identity) GetCheckImperfectTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("check-imperfect-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Identity) SetCheckImperfectTimestamp(value bool) error {
	return e.Element.SetProperty("check-imperfect-timestamp", value)
}

// datarate (int)
//
// (Re)timestamps buffers with number of bytes per second (0 = inactive)

func (e *Identity) GetDatarate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("datarate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Identity) SetDatarate(value int) error {
	return e.Element.SetProperty("datarate", value)
}

// drop-allocation (bool)
//
// Don't forward allocation queries

func (e *Identity) GetDropAllocation() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-allocation")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Identity) SetDropAllocation(value bool) error {
	return e.Element.SetProperty("drop-allocation", value)
}

// drop-buffer-flags (GstBufferFlags)
//
// Drop buffers with the given flags.

func (e *Identity) GetDropBufferFlags() (interface{}, error) {
	return e.Element.GetProperty("drop-buffer-flags")
}

func (e *Identity) SetDropBufferFlags(value interface{}) error {
	return e.Element.SetProperty("drop-buffer-flags", value)
}

// drop-probability (float32)
//
// The Probability a buffer is dropped

func (e *Identity) GetDropProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("drop-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Identity) SetDropProbability(value float32) error {
	return e.Element.SetProperty("drop-probability", value)
}

// dump (bool)
//
// Dump buffer contents to stdout

func (e *Identity) GetDump() (bool, error) {
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

func (e *Identity) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

// eos-after (int)
//
// EOS after N buffers.

func (e *Identity) GetEosAfter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("eos-after")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Identity) SetEosAfter(value int) error {
	return e.Element.SetProperty("eos-after", value)
}

// error-after (int)
//
// Error after N buffers

func (e *Identity) GetErrorAfter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("error-after")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Identity) SetErrorAfter(value int) error {
	return e.Element.SetProperty("error-after", value)
}

// last-message (string)
//
// last-message

func (e *Identity) GetLastMessage() (string, error) {
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

// signal-handoffs (bool)
//
// If set to TRUE, the identity will emit a handoff signal when handling a buffer.
// When set to FALSE, no signal will be emitted, which might improve performance.

func (e *Identity) GetSignalHandoffs() (bool, error) {
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

func (e *Identity) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

// silent (bool)
//
// silent

func (e *Identity) GetSilent() (bool, error) {
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

func (e *Identity) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// single-segment (bool)
//
// Timestamp buffers and eat segments so as to appear as one segment

func (e *Identity) GetSingleSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("single-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Identity) SetSingleSegment(value bool) error {
	return e.Element.SetProperty("single-segment", value)
}

// sleep-time (uint)
//
// Microseconds to sleep between processing

func (e *Identity) GetSleepTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sleep-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Identity) SetSleepTime(value uint) error {
	return e.Element.SetProperty("sleep-time", value)
}

// stats (GstStructure)
//
// Various statistics. This property returns a GstStructure
// with name application/x-identity-stats with the following fields:

// sync (bool)
//
// Synchronize to pipeline clock

func (e *Identity) GetSync() (bool, error) {
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

func (e *Identity) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds for synchronisation, negative for earlier sync

func (e *Identity) GetTsOffset() (int64, error) {
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

func (e *Identity) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

