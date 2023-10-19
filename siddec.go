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

type Siddec struct {
	Element *gst.Element
}

func NewSiddec() (*Siddec, error) {
	e, err := gst.NewElement("siddec")
	if err != nil {
		return nil, err
	}
	return &Siddec{Element: e}, nil
}

func NewSiddecWithName(name string) (*Siddec, error) {
	e, err := gst.NewElementWithName("siddec", name)
	if err != nil {
		return nil, err
	}
	return &Siddec{Element: e}, nil
}

// ----- Properties -----

// blocksize (uint)
//
// Size in bytes to output per buffer

func (e *Siddec) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Siddec) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// clock (GstSidClock)
//
// clock

func (e *Siddec) GetClock() (GstSidClock, error) {
	var value GstSidClock
	var ok bool
	v, err := e.Element.GetProperty("clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSidClock)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSidClock")
	}
	return value, nil
}

func (e *Siddec) SetClock(value GstSidClock) error {
	e.Element.SetArg("clock", string(value))
	return nil
}

// filter (bool)
//
// filter

func (e *Siddec) GetFilter() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Siddec) SetFilter(value bool) error {
	return e.Element.SetProperty("filter", value)
}

// force-speed (bool)
//
// force_speed

func (e *Siddec) GetForceSpeed() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Siddec) SetForceSpeed(value bool) error {
	return e.Element.SetProperty("force-speed", value)
}

// measured-volume (bool)
//
// measured_volume

func (e *Siddec) GetMeasuredVolume() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("measured-volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Siddec) SetMeasuredVolume(value bool) error {
	return e.Element.SetProperty("measured-volume", value)
}

// memory (GstSidMemory)
//
// memory

func (e *Siddec) GetMemory() (GstSidMemory, error) {
	var value GstSidMemory
	var ok bool
	v, err := e.Element.GetProperty("memory")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSidMemory)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSidMemory")
	}
	return value, nil
}

func (e *Siddec) SetMemory(value GstSidMemory) error {
	e.Element.SetArg("memory", string(value))
	return nil
}

// metadata (GstCaps)
//
// Metadata

func (e *Siddec) GetMetadata() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("metadata")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// mos8580 (bool)
//
// mos8580

func (e *Siddec) GetMos8580() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mos8580")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Siddec) SetMos8580(value bool) error {
	return e.Element.SetProperty("mos8580", value)
}

// tune (int)
//
// tune

func (e *Siddec) GetTune() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("tune")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Siddec) SetTune(value int) error {
	return e.Element.SetProperty("tune", value)
}

// ----- Constants -----

type GstSidClock string

const (
	GstSidClockPal GstSidClock = "pal" // pal (1) – PAL
	GstSidClockNtsc GstSidClock = "ntsc" // ntsc (2) – NTSC
)

type GstSidMemory string

const (
	GstSidMemoryBankSwitching GstSidMemory = "bank-switching" // bank-switching (32) – Bank Switching
	GstSidMemoryTransparentRom GstSidMemory = "transparent-rom" // transparent-rom (33) – Transparent ROM
	GstSidMemoryPlaysidEnvironment GstSidMemory = "playsid-environment" // playsid-environment (34) – Playsid Environment
)

