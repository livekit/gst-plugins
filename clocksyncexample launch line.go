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

type ClocksyncexampleLaunchLine struct {
	Element *gst.Element
}

func NewClocksyncexampleLaunchLine() (*ClocksyncexampleLaunchLine, error) {
	e, err := gst.NewElement("clocksyncExample launch line")
	if err != nil {
		return nil, err
	}
	return &ClocksyncexampleLaunchLine{Element: e}, nil
}

func NewClocksyncexampleLaunchLineWithName(name string) (*ClocksyncexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("clocksyncExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &ClocksyncexampleLaunchLine{Element: e}, nil
}

// ----- Properties -----

// qos (bool)
//
// Generate Quality-of-Service events upstream.

func (e *ClocksyncexampleLaunchLine) GetQos() (bool, error) {
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

func (e *ClocksyncexampleLaunchLine) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// sync (bool)
//
// Synchronize to pipeline clock

func (e *ClocksyncexampleLaunchLine) GetSync() (bool, error) {
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

func (e *ClocksyncexampleLaunchLine) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// sync-to-first (bool)
//
// When enabled, clocksync elemenet will adjust "ts-offset" value
// automatically by using given timestamp of the first buffer and running
// time of pipeline, so that clocksync element can output the first buffer
// immediately without clock waiting.

func (e *ClocksyncexampleLaunchLine) GetSyncToFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync-to-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *ClocksyncexampleLaunchLine) SetSyncToFirst(value bool) error {
	return e.Element.SetProperty("sync-to-first", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds for synchronisation, negative for earlier sync

func (e *ClocksyncexampleLaunchLine) GetTsOffset() (int64, error) {
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

func (e *ClocksyncexampleLaunchLine) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

