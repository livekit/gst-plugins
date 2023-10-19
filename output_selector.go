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

type OutputSelector struct {
	Element *gst.Element
}

func NewOutputSelector() (*OutputSelector, error) {
	e, err := gst.NewElement("output-selector")
	if err != nil {
		return nil, err
	}
	return &OutputSelector{Element: e}, nil
}

func NewOutputSelectorWithName(name string) (*OutputSelector, error) {
	e, err := gst.NewElementWithName("output-selector", name)
	if err != nil {
		return nil, err
	}
	return &OutputSelector{Element: e}, nil
}

// ----- Properties -----

// active-pad (GstPad)
//
// Currently active src pad

func (e *OutputSelector) GetActivePad() (interface{}, error) {
	return e.Element.GetProperty("active-pad")
}

func (e *OutputSelector) SetActivePad(value interface{}) error {
	return e.Element.SetProperty("active-pad", value)
}

// pad-negotiation-mode (GstOutputSelectorPadNegotiationMode)
//
// The mode to be used for pad negotiation

func (e *OutputSelector) GetPadNegotiationMode() (GstOutputSelectorPadNegotiationMode, error) {
	var value GstOutputSelectorPadNegotiationMode
	var ok bool
	v, err := e.Element.GetProperty("pad-negotiation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOutputSelectorPadNegotiationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOutputSelectorPadNegotiationMode")
	}
	return value, nil
}

func (e *OutputSelector) SetPadNegotiationMode(value GstOutputSelectorPadNegotiationMode) error {
	e.Element.SetArg("pad-negotiation-mode", string(value))
	return nil
}

// resend-latest (bool)
//
// Resend latest buffer after a switch to a new pad

func (e *OutputSelector) GetResendLatest() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("resend-latest")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *OutputSelector) SetResendLatest(value bool) error {
	return e.Element.SetProperty("resend-latest", value)
}

// ----- Constants -----

type GstOutputSelectorPadNegotiationMode string

const (
	GstOutputSelectorPadNegotiationModeNone GstOutputSelectorPadNegotiationMode = "none" // none (0) – None
	GstOutputSelectorPadNegotiationModeAll GstOutputSelectorPadNegotiationMode = "all" // all (1) – All
	GstOutputSelectorPadNegotiationModeActive GstOutputSelectorPadNegotiationMode = "active" // active (2) – Active
)

