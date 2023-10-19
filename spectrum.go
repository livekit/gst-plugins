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

type Spectrum struct {
	*base.GstBaseTransform
}

func NewSpectrum() (*Spectrum, error) {
	e, err := gst.NewElement("spectrum")
	if err != nil {
		return nil, err
	}
	return &Spectrum{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewSpectrumWithName(name string) (*Spectrum, error) {
	e, err := gst.NewElementWithName("spectrum", name)
	if err != nil {
		return nil, err
	}
	return &Spectrum{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bands (uint)
//
// Number of frequency bands

func (e *Spectrum) GetBands() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bands")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Spectrum) SetBands(value uint) error {
	return e.Element.SetProperty("bands", value)
}

// interval (uint64)
//
// Interval of time between message posts (in nanoseconds)

func (e *Spectrum) GetInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Spectrum) SetInterval(value uint64) error {
	return e.Element.SetProperty("interval", value)
}

// message-magnitude (bool)
//
// Whether to add a 'magnitude' field to the structure of any 'spectrum' element messages posted on the bus

func (e *Spectrum) GetMessageMagnitude() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message-magnitude")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Spectrum) SetMessageMagnitude(value bool) error {
	return e.Element.SetProperty("message-magnitude", value)
}

// message-phase (bool)
//
// Whether to add a 'phase' field to the structure of any 'spectrum' element messages posted on the bus

func (e *Spectrum) GetMessagePhase() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message-phase")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Spectrum) SetMessagePhase(value bool) error {
	return e.Element.SetProperty("message-phase", value)
}

// multi-channel (bool)
//
// Send separate results for each channel

func (e *Spectrum) GetMultiChannel() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("multi-channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Spectrum) SetMultiChannel(value bool) error {
	return e.Element.SetProperty("multi-channel", value)
}

// post-messages (bool)
//
// Whether to post a 'spectrum' element message on the bus for each passed interval

func (e *Spectrum) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Spectrum) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

// threshold (int)
//
// dB threshold for result. All lower values will be set to this

func (e *Spectrum) GetThreshold() (int, error) {
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

func (e *Spectrum) SetThreshold(value int) error {
	return e.Element.SetProperty("threshold", value)
}

