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

type Fluiddec struct {
	Element *gst.Element
}

func NewFluiddec() (*Fluiddec, error) {
	e, err := gst.NewElement("fluiddec")
	if err != nil {
		return nil, err
	}
	return &Fluiddec{Element: e}, nil
}

func NewFluiddecWithName(name string) (*Fluiddec, error) {
	e, err := gst.NewElementWithName("fluiddec", name)
	if err != nil {
		return nil, err
	}
	return &Fluiddec{Element: e}, nil
}

// ----- Properties -----

// soundfont (string)
//
// the filename of a soundfont (NULL for default)

func (e *Fluiddec) GetSoundfont() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("soundfont")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Fluiddec) SetSoundfont(value string) error {
	return e.Element.SetProperty("soundfont", value)
}

// synth-chorus (bool)
//
// Turn the chorus on or off

func (e *Fluiddec) GetSynthChorus() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synth-chorus")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fluiddec) SetSynthChorus(value bool) error {
	return e.Element.SetProperty("synth-chorus", value)
}

// synth-gain (float64)
//
// Set the master gain

func (e *Fluiddec) GetSynthGain() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("synth-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Fluiddec) SetSynthGain(value float64) error {
	return e.Element.SetProperty("synth-gain", value)
}

// synth-polyphony (int)
//
// The number of simultaneous voices

func (e *Fluiddec) GetSynthPolyphony() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("synth-polyphony")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fluiddec) SetSynthPolyphony(value int) error {
	return e.Element.SetProperty("synth-polyphony", value)
}

// synth-reverb (bool)
//
// Turn the reverb on or off

func (e *Fluiddec) GetSynthReverb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("synth-reverb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fluiddec) SetSynthReverb(value bool) error {
	return e.Element.SetProperty("synth-reverb", value)
}

