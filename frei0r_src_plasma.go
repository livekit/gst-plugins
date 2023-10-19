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

type Frei0RSrcPlasma struct {
	*base.GstPushSrc
}

func NewFrei0RSrcPlasma() (*Frei0RSrcPlasma, error) {
	e, err := gst.NewElement("frei0r-src-plasma")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcPlasma{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcPlasmaWithName(name string) (*Frei0RSrcPlasma, error) {
	e, err := gst.NewElementWithName("frei0r-src-plasma", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcPlasma{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// param-1-move (float64)

func (e *Frei0RSrcPlasma) GetParam1Move() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-1-move")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPlasma) SetParam1Move(value float64) error {
	return e.Element.SetProperty("param-1-move", value)
}

// param-1-speed (float64)

func (e *Frei0RSrcPlasma) GetParam1Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-1-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPlasma) SetParam1Speed(value float64) error {
	return e.Element.SetProperty("param-1-speed", value)
}

// param-2-move (float64)

func (e *Frei0RSrcPlasma) GetParam2Move() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-2-move")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPlasma) SetParam2Move(value float64) error {
	return e.Element.SetProperty("param-2-move", value)
}

// param-2-speed (float64)

func (e *Frei0RSrcPlasma) GetParam2Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-2-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPlasma) SetParam2Speed(value float64) error {
	return e.Element.SetProperty("param-2-speed", value)
}

// param-3-speed (float64)

func (e *Frei0RSrcPlasma) GetParam3Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-3-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPlasma) SetParam3Speed(value float64) error {
	return e.Element.SetProperty("param-3-speed", value)
}

// param-4-speed (float64)

func (e *Frei0RSrcPlasma) GetParam4Speed() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("param-4-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPlasma) SetParam4Speed(value float64) error {
	return e.Element.SetProperty("param-4-speed", value)
}

