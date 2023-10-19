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

type Modplug struct {
	Element *gst.Element
}

func NewModplug() (*Modplug, error) {
	e, err := gst.NewElement("modplug")
	if err != nil {
		return nil, err
	}
	return &Modplug{Element: e}, nil
}

func NewModplugWithName(name string) (*Modplug, error) {
	e, err := gst.NewElementWithName("modplug", name)
	if err != nil {
		return nil, err
	}
	return &Modplug{Element: e}, nil
}

// ----- Properties -----

// megabass (bool)
//
// Megabass

func (e *Modplug) GetMegabass() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("megabass")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Modplug) SetMegabass(value bool) error {
	return e.Element.SetProperty("megabass", value)
}

// megabass-amount (int)
//
// Megabass amount

func (e *Modplug) GetMegabassAmount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("megabass-amount")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Modplug) SetMegabassAmount(value int) error {
	return e.Element.SetProperty("megabass-amount", value)
}

// megabass-range (int)
//
// Megabass range

func (e *Modplug) GetMegabassRange() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("megabass-range")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Modplug) SetMegabassRange(value int) error {
	return e.Element.SetProperty("megabass-range", value)
}

// noise-reduction (bool)
//
// noise reduction

func (e *Modplug) GetNoiseReduction() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("noise-reduction")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Modplug) SetNoiseReduction(value bool) error {
	return e.Element.SetProperty("noise-reduction", value)
}

// oversamp (bool)
//
// oversamp

func (e *Modplug) GetOversamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("oversamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Modplug) SetOversamp(value bool) error {
	return e.Element.SetProperty("oversamp", value)
}

// reverb (bool)
//
// Reverb

func (e *Modplug) GetReverb() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reverb")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Modplug) SetReverb(value bool) error {
	return e.Element.SetProperty("reverb", value)
}

// reverb-delay (int)
//
// Reverb delay

func (e *Modplug) GetReverbDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("reverb-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Modplug) SetReverbDelay(value int) error {
	return e.Element.SetProperty("reverb-delay", value)
}

// reverb-depth (int)
//
// Reverb depth

func (e *Modplug) GetReverbDepth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("reverb-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Modplug) SetReverbDepth(value int) error {
	return e.Element.SetProperty("reverb-depth", value)
}

// songname (string)
//
// The song name

func (e *Modplug) GetSongname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("songname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// surround (bool)
//
// Surround

func (e *Modplug) GetSurround() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("surround")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Modplug) SetSurround(value bool) error {
	return e.Element.SetProperty("surround", value)
}

// surround-delay (int)
//
// Surround delay

func (e *Modplug) GetSurroundDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("surround-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Modplug) SetSurroundDelay(value int) error {
	return e.Element.SetProperty("surround-delay", value)
}

// surround-depth (int)
//
// Surround depth

func (e *Modplug) GetSurroundDepth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("surround-depth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Modplug) SetSurroundDepth(value int) error {
	return e.Element.SetProperty("surround-depth", value)
}

