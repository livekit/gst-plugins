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

type Decklinkaudiosink struct {
	*base.GstBaseSink
}

func NewDecklinkaudiosink() (*Decklinkaudiosink, error) {
	e, err := gst.NewElement("decklinkaudiosink")
	if err != nil {
		return nil, err
	}
	return &Decklinkaudiosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDecklinkaudiosinkWithName(name string) (*Decklinkaudiosink, error) {
	e, err := gst.NewElementWithName("decklinkaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &Decklinkaudiosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// alignment-threshold (uint64)
//
// Timestamp alignment threshold in nanoseconds

func (e *Decklinkaudiosink) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decklinkaudiosink) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

// buffer-time (uint64)
//
// Size of audio buffer in microseconds, this is the minimum latency that the sink reports

func (e *Decklinkaudiosink) GetBufferTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decklinkaudiosink) SetBufferTime(value uint64) error {
	return e.Element.SetProperty("buffer-time", value)
}

// device-number (int)
//
// Output device instance to use

func (e *Decklinkaudiosink) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decklinkaudiosink) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

// discont-wait (uint64)
//
// Window of time in nanoseconds to wait before creating a discontinuity

func (e *Decklinkaudiosink) GetDiscontWait() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("discont-wait")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decklinkaudiosink) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

// hw-serial-number (string)
//
// The serial number (hardware ID) of the Decklink card

func (e *Decklinkaudiosink) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// persistent-id (int64)
//
// Decklink device to use. Higher priority than "device-number".
// BMDDeckLinkPersistentID is a device speciﬁc, 32-bit unique identiﬁer.
// It is stable even when the device is plugged in a diļ¬erent connector,
// across reboots, and when plugged into diļ¬erent computers.

func (e *Decklinkaudiosink) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Decklinkaudiosink) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

