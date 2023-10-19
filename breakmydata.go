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

	"github.com/livekit/gstplugins/base"
)

type Breakmydata struct {
	*base.GstBaseTransform
}

func NewBreakmydata() (*Breakmydata, error) {
	e, err := gst.NewElement("breakmydata")
	if err != nil {
		return nil, err
	}
	return &Breakmydata{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewBreakmydataWithName(name string) (*Breakmydata, error) {
	e, err := gst.NewElementWithName("breakmydata", name)
	if err != nil {
		return nil, err
	}
	return &Breakmydata{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// probability (float64)
//
// probability for each byte in the buffer to be changed

func (e *Breakmydata) GetProbability() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Breakmydata) SetProbability(value float64) error {
	return e.Element.SetProperty("probability", value)
}

// seed (uint)
//
// seed for randomness (initialized when going from READY to PAUSED)

func (e *Breakmydata) GetSeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Breakmydata) SetSeed(value uint) error {
	return e.Element.SetProperty("seed", value)
}

// set-to (int)
//
// set changed bytes to this value (-1 means random value

func (e *Breakmydata) GetSetTo() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("set-to")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Breakmydata) SetSetTo(value int) error {
	return e.Element.SetProperty("set-to", value)
}

// skip (uint)
//
// amount of bytes skipped at the beginning of stream

func (e *Breakmydata) GetSkip() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Breakmydata) SetSkip(value uint) error {
	return e.Element.SetProperty("skip", value)
}

