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

type Tsdemux struct {
	Element *gst.Element
}

func NewTsdemux() (*Tsdemux, error) {
	e, err := gst.NewElement("tsdemux")
	if err != nil {
		return nil, err
	}
	return &Tsdemux{Element: e}, nil
}

func NewTsdemuxWithName(name string) (*Tsdemux, error) {
	e, err := gst.NewElementWithName("tsdemux", name)
	if err != nil {
		return nil, err
	}
	return &Tsdemux{Element: e}, nil
}

// ----- Properties -----

// emit-stats (bool)
//
// Emit messages for every pcr/opcr/pts/dts

func (e *Tsdemux) GetEmitStats() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("emit-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tsdemux) SetEmitStats(value bool) error {
	return e.Element.SetProperty("emit-stats", value)
}

// latency (int)
//
// Latency to add for smooth demuxing (in ms)

func (e *Tsdemux) GetLatency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tsdemux) SetLatency(value int) error {
	return e.Element.SetProperty("latency", value)
}

// program-number (int)
//
// Program Number to demux for (-1 to ignore)

func (e *Tsdemux) GetProgramNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("program-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tsdemux) SetProgramNumber(value int) error {
	return e.Element.SetProperty("program-number", value)
}

// send-scte35-events (bool)
//
// Whether SCTE 35 sections should be forwarded as events.

