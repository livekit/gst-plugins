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

type Opusdec struct {
	Element *gst.Element
}

func NewOpusdec() (*Opusdec, error) {
	e, err := gst.NewElement("opusdec")
	if err != nil {
		return nil, err
	}
	return &Opusdec{Element: e}, nil
}

func NewOpusdecWithName(name string) (*Opusdec, error) {
	e, err := gst.NewElementWithName("opusdec", name)
	if err != nil {
		return nil, err
	}
	return &Opusdec{Element: e}, nil
}

// ----- Properties -----

// apply-gain (bool)
//
// Apply gain if any is specified in the header

func (e *Opusdec) GetApplyGain() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("apply-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Opusdec) SetApplyGain(value bool) error {
	return e.Element.SetProperty("apply-gain", value)
}

// phase-inversion (bool)
//
// Set to true to enable phase inversion, this will slightly improve stereo quality, but will have side effects when downmixed to mono.

func (e *Opusdec) GetPhaseInversion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("phase-inversion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Opusdec) SetPhaseInversion(value bool) error {
	return e.Element.SetProperty("phase-inversion", value)
}

// stats (GstStructure)
//
// Various decoder statistics. This property returns a GstStructure
// with name application/x-opusdec-stats with the following fields:

// use-inband-fec (bool)
//
// Use forward error correction if available (needs PLC enabled)

func (e *Opusdec) GetUseInbandFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-inband-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Opusdec) SetUseInbandFec(value bool) error {
	return e.Element.SetProperty("use-inband-fec", value)
}

