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

type Audioecho struct {
	*base.GstBaseTransform
}

func NewAudioecho() (*Audioecho, error) {
	e, err := gst.NewElement("audioecho")
	if err != nil {
		return nil, err
	}
	return &Audioecho{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudioechoWithName(name string) (*Audioecho, error) {
	e, err := gst.NewElementWithName("audioecho", name)
	if err != nil {
		return nil, err
	}
	return &Audioecho{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// delay (uint64)
//
// Delay of the echo in nanoseconds

func (e *Audioecho) GetDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audioecho) SetDelay(value uint64) error {
	return e.Element.SetProperty("delay", value)
}

// feedback (float32)
//
// Amount of feedback

func (e *Audioecho) GetFeedback() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("feedback")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audioecho) SetFeedback(value float32) error {
	return e.Element.SetProperty("feedback", value)
}

// intensity (float32)
//
// Intensity of the echo

func (e *Audioecho) GetIntensity() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("intensity")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Audioecho) SetIntensity(value float32) error {
	return e.Element.SetProperty("intensity", value)
}

// max-delay (uint64)
//
// Maximum delay of the echo in nanoseconds (can't be changed in PLAYING or PAUSED state)

func (e *Audioecho) GetMaxDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audioecho) SetMaxDelay(value uint64) error {
	return e.Element.SetProperty("max-delay", value)
}

// surround-delay (bool)
//
// Delay Surround Channels when TRUE instead of applying an echo effect

func (e *Audioecho) GetSurroundDelay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("surround-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audioecho) SetSurroundDelay(value bool) error {
	return e.Element.SetProperty("surround-delay", value)
}

// surround-mask (uint64)
//
// A bitmask of channels that are considered surround and delayed when surround-delay = TRUE

func (e *Audioecho) GetSurroundMask() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("surround-mask")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audioecho) SetSurroundMask(value uint64) error {
	return e.Element.SetProperty("surround-mask", value)
}

