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

type Tee struct {
	Element *gst.Element
}

func NewTee() (*Tee, error) {
	e, err := gst.NewElement("tee")
	if err != nil {
		return nil, err
	}
	return &Tee{Element: e}, nil
}

func NewTeeWithName(name string) (*Tee, error) {
	e, err := gst.NewElementWithName("tee", name)
	if err != nil {
		return nil, err
	}
	return &Tee{Element: e}, nil
}

// ----- Properties -----

// alloc-pad (GstPad)
//
// The pad ALLOCATION queries will be proxied to (DEPRECATED, has no effect)

func (e *Tee) GetAllocPad() (interface{}, error) {
	return e.Element.GetProperty("alloc-pad")
}

func (e *Tee) SetAllocPad(value interface{}) error {
	return e.Element.SetProperty("alloc-pad", value)
}

// allow-not-linked (bool)
//
// This property makes sink pad return GST_FLOW_OK even if there are no
// source pads or any of them is linked.

// has-chain (bool)
//
// If the element can operate in push mode

func (e *Tee) GetHasChain() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("has-chain")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tee) SetHasChain(value bool) error {
	return e.Element.SetProperty("has-chain", value)
}

// last-message (string)
//
// The message describing current status

func (e *Tee) GetLastMessage() (string, error) {
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

// num-src-pads (int)
//
// The number of source pads

func (e *Tee) GetNumSrcPads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-src-pads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// pull-mode (GstTeePullMode)
//
// Behavior of tee in pull mode

func (e *Tee) GetPullMode() (GstTeePullMode, error) {
	var value GstTeePullMode
	var ok bool
	v, err := e.Element.GetProperty("pull-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTeePullMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTeePullMode")
	}
	return value, nil
}

func (e *Tee) SetPullMode(value GstTeePullMode) error {
	e.Element.SetArg("pull-mode", string(value))
	return nil
}

// silent (bool)
//
// Don't produce last_message events

func (e *Tee) GetSilent() (bool, error) {
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

func (e *Tee) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// ----- Constants -----

type GstTeePullMode string

const (
	GstTeePullModeNever GstTeePullMode = "never" // never (0) – Never activate in pull mode
	GstTeePullModeSingle GstTeePullMode = "single" // single (1) – Only one src pad can be active in pull mode
)

