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

type Pitch struct {
	Element *gst.Element
}

func NewPitch() (*Pitch, error) {
	e, err := gst.NewElement("pitch")
	if err != nil {
		return nil, err
	}
	return &Pitch{Element: e}, nil
}

func NewPitchWithName(name string) (*Pitch, error) {
	e, err := gst.NewElementWithName("pitch", name)
	if err != nil {
		return nil, err
	}
	return &Pitch{Element: e}, nil
}

// ----- Properties -----

// output-rate (float32)
//
// Output rate on downstream segment events

func (e *Pitch) GetOutputRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("output-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Pitch) SetOutputRate(value float32) error {
	return e.Element.SetProperty("output-rate", value)
}

// pitch (float32)
//
// Audio stream pitch

func (e *Pitch) GetPitch() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pitch")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Pitch) SetPitch(value float32) error {
	return e.Element.SetProperty("pitch", value)
}

// rate (float32)
//
// Audio stream rate

func (e *Pitch) GetRate() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Pitch) SetRate(value float32) error {
	return e.Element.SetProperty("rate", value)
}

// tempo (float32)
//
// Audio stream tempo

func (e *Pitch) GetTempo() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tempo")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Pitch) SetTempo(value float32) error {
	return e.Element.SetProperty("tempo", value)
}

