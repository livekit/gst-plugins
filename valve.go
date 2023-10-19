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

type Valve struct {
	Element *gst.Element
}

func NewValve() (*Valve, error) {
	e, err := gst.NewElement("valve")
	if err != nil {
		return nil, err
	}
	return &Valve{Element: e}, nil
}

func NewValveWithName(name string) (*Valve, error) {
	e, err := gst.NewElementWithName("valve", name)
	if err != nil {
		return nil, err
	}
	return &Valve{Element: e}, nil
}

// ----- Properties -----

// drop (bool)
//
// Whether to drop buffers and events or let them through

func (e *Valve) GetDrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Valve) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

// drop-mode (GstValveDropMode)
//
// Drop mode to use. By default all buffers and events are dropped.

func (e *Valve) GetDropMode() (GstValveDropMode, error) {
	var value GstValveDropMode
	var ok bool
	v, err := e.Element.GetProperty("drop-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstValveDropMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstValveDropMode")
	}
	return value, nil
}

func (e *Valve) SetDropMode(value GstValveDropMode) error {
	e.Element.SetArg("drop-mode", string(value))
	return nil
}

// ----- Constants -----

type GstValveDropMode string

const (
	GstValveDropModeDropAll GstValveDropMode = "drop-all" // drop-all (0) – Drop all buffers and events
	GstValveDropModeForwardStickyEvents GstValveDropMode = "forward-sticky-events" // forward-sticky-events (1) – Drop all buffers but forward sticky events
	GstValveDropModeTransformToGap GstValveDropMode = "transform-to-gap" // transform-to-gap (2) – Convert all dropped buffers into gap events and forward sticky events
)

