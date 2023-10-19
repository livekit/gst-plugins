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

type InputSelector struct {
	Element *gst.Element
}

func NewInputSelector() (*InputSelector, error) {
	e, err := gst.NewElement("input-selector")
	if err != nil {
		return nil, err
	}
	return &InputSelector{Element: e}, nil
}

func NewInputSelectorWithName(name string) (*InputSelector, error) {
	e, err := gst.NewElementWithName("input-selector", name)
	if err != nil {
		return nil, err
	}
	return &InputSelector{Element: e}, nil
}

// ----- Properties -----

// active-pad (GstPad)
//
// The currently active sink pad

func (e *InputSelector) GetActivePad() (interface{}, error) {
	return e.Element.GetProperty("active-pad")
}

func (e *InputSelector) SetActivePad(value interface{}) error {
	return e.Element.SetProperty("active-pad", value)
}

// cache-buffers (bool)
//
// If set to TRUE and GstInputSelector:sync-streams is also set to TRUE,
// the active pad will cache the buffers still considered valid (after current
// running time, see sync-mode) to avoid missing frames if/when the pad is
// reactivated.

// drop-backwards (bool)
//
// If set to TRUE and GstInputSelector:sync-streams is also set to TRUE,
// every time the input is switched, buffers that would go backwards related
// to the last output buffer pre-switch will be dropped.

func (e *InputSelector) GetDropBackwards() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-backwards")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *InputSelector) SetDropBackwards(value bool) error {
	return e.Element.SetProperty("drop-backwards", value)
}

// n-pads (uint)
//
// The number of sink pads

func (e *InputSelector) GetNPads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("n-pads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// sync-mode (GstInputSelectorSyncMode)
//
// Select how input-selector will sync buffers when in sync-streams mode.

// sync-streams (bool)
//
// If set to TRUE all inactive streams will be synced to the
// running time of the active stream or to the current clock.

// ----- Constants -----

type GstInputSelectorSyncMode string

const (
	GstInputSelectorSyncModeActiveSegment GstInputSelectorSyncMode = "active-segment" // active-segment (0) – Sync using the current active segment
	GstInputSelectorSyncModeClock GstInputSelectorSyncMode = "clock" // clock (1) – Sync using the clock
)

