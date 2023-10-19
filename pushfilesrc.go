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

type Pushfilesrc struct {
	Element *gst.Element
}

func NewPushfilesrc() (*Pushfilesrc, error) {
	e, err := gst.NewElement("pushfilesrc")
	if err != nil {
		return nil, err
	}
	return &Pushfilesrc{Element: e}, nil
}

func NewPushfilesrcWithName(name string) (*Pushfilesrc, error) {
	e, err := gst.NewElementWithName("pushfilesrc", name)
	if err != nil {
		return nil, err
	}
	return &Pushfilesrc{Element: e}, nil
}

// ----- Properties -----

// applied-rate (float64)
//
// Applied rate to use in TIME SEGMENT

func (e *Pushfilesrc) GetAppliedRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("applied-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Pushfilesrc) SetAppliedRate(value float64) error {
	return e.Element.SetProperty("applied-rate", value)
}

// initial-timestamp (uint64)
//
// Initial Buffer Timestamp (if time-segment TRUE)

func (e *Pushfilesrc) GetInitialTimestamp() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("initial-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Pushfilesrc) SetInitialTimestamp(value uint64) error {
	return e.Element.SetProperty("initial-timestamp", value)
}

// location (string)
//
// Location of the file to read

func (e *Pushfilesrc) GetLocation() (string, error) {
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

func (e *Pushfilesrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// rate (float64)
//
// Rate to use in TIME SEGMENT

func (e *Pushfilesrc) GetRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Pushfilesrc) SetRate(value float64) error {
	return e.Element.SetProperty("rate", value)
}

// start-time (int64)
//
// Initial Start Time (if time-segment TRUE)

func (e *Pushfilesrc) GetStartTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("start-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Pushfilesrc) SetStartTime(value int64) error {
	return e.Element.SetProperty("start-time", value)
}

// stream-time (int64)
//
// Initial Stream Time (if time-segment TRUE)

func (e *Pushfilesrc) GetStreamTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("stream-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Pushfilesrc) SetStreamTime(value int64) error {
	return e.Element.SetProperty("stream-time", value)
}

// time-segment (bool)
//
// Emit TIME SEGMENTS

func (e *Pushfilesrc) GetTimeSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("time-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Pushfilesrc) SetTimeSegment(value bool) error {
	return e.Element.SetProperty("time-segment", value)
}

