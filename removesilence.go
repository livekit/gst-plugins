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

type Removesilence struct {
	*base.GstBaseTransform
}

func NewRemovesilence() (*Removesilence, error) {
	e, err := gst.NewElement("removesilence")
	if err != nil {
		return nil, err
	}
	return &Removesilence{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRemovesilenceWithName(name string) (*Removesilence, error) {
	e, err := gst.NewElementWithName("removesilence", name)
	if err != nil {
		return nil, err
	}
	return &Removesilence{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// hysteresis (uint64)
//
// Set the hysteresis (on samples) used on the internal VAD

func (e *Removesilence) GetHysteresis() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("hysteresis")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Removesilence) SetHysteresis(value uint64) error {
	return e.Element.SetProperty("hysteresis", value)
}

// minimum-silence-buffers (uint)
//
// Define the minimum number of consecutive silence buffers before removing silence, 0 means disabled. This will not introduce latency

func (e *Removesilence) GetMinimumSilenceBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("minimum-silence-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Removesilence) SetMinimumSilenceBuffers(value uint) error {
	return e.Element.SetProperty("minimum-silence-buffers", value)
}

// minimum-silence-time (uint64)
//
// Define the minimum silence time in nanoseconds before removing  silence, 0 means disabled. This will not introduce latency

func (e *Removesilence) GetMinimumSilenceTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("minimum-silence-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Removesilence) SetMinimumSilenceTime(value uint64) error {
	return e.Element.SetProperty("minimum-silence-time", value)
}

// remove (bool)
//
// Set to true to remove silence from the stream, false otherwise

func (e *Removesilence) GetRemove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Removesilence) SetRemove(value bool) error {
	return e.Element.SetProperty("remove", value)
}

// silent (bool)
//
// Disable/enable bus message notifications for silence detected/finished

func (e *Removesilence) GetSilent() (bool, error) {
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

func (e *Removesilence) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// squash (bool)
//
// Set to true to retimestamp buffers when silence is removed and so avoid timestamp gap

func (e *Removesilence) GetSquash() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("squash")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Removesilence) SetSquash(value bool) error {
	return e.Element.SetProperty("squash", value)
}

// threshold (int)
//
// Set the silence threshold used on the internal VAD in dB

func (e *Removesilence) GetThreshold() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Removesilence) SetThreshold(value int) error {
	return e.Element.SetProperty("threshold", value)
}

