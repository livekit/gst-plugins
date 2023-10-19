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

type Tonegeneratesrc struct {
	*base.GstPushSrc
}

func NewTonegeneratesrc() (*Tonegeneratesrc, error) {
	e, err := gst.NewElement("tonegeneratesrc")
	if err != nil {
		return nil, err
	}
	return &Tonegeneratesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewTonegeneratesrcWithName(name string) (*Tonegeneratesrc, error) {
	e, err := gst.NewElementWithName("tonegeneratesrc", name)
	if err != nil {
		return nil, err
	}
	return &Tonegeneratesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// blocksize (uint)
//
// Size in bytes to read per buffer (-1 = default)

func (e *Tonegeneratesrc) GetBlocksize() (uint, error) {
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

func (e *Tonegeneratesrc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// do-timestamp (bool)
//
// Apply current stream time to buffers

func (e *Tonegeneratesrc) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

// freq (int)
//
// Frequency of test signal

func (e *Tonegeneratesrc) GetFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetFreq(value int) error {
	return e.Element.SetProperty("freq", value)
}

// freq2 (int)
//
// Frequency of second telephony tone component

func (e *Tonegeneratesrc) GetFreq2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("freq2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetFreq2(value int) error {
	return e.Element.SetProperty("freq2", value)
}

// num-buffers (int)
//
// Number of buffers to output before sending EOS (-1 = unlimited)

func (e *Tonegeneratesrc) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// off-time (int)
//
// Time of the first period  when the tone signal is off

func (e *Tonegeneratesrc) GetOffTime() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("off-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetOffTime(value int) error {
	return e.Element.SetProperty("off-time", value)
}

// off-time2 (int)
//
// Time of the second period  when the tone signal is off

func (e *Tonegeneratesrc) GetOffTime2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("off-time2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetOffTime2(value int) error {
	return e.Element.SetProperty("off-time2", value)
}

// on-time (int)
//
// Time of the first period  when the tone signal is present

func (e *Tonegeneratesrc) GetOnTime() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("on-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetOnTime(value int) error {
	return e.Element.SetProperty("on-time", value)
}

// on-time2 (int)
//
// Time of the second period  when the tone signal is present

func (e *Tonegeneratesrc) GetOnTime2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("on-time2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetOnTime2(value int) error {
	return e.Element.SetProperty("on-time2", value)
}

// repeat (bool)
//
// Whether to repeat specified tone indefinitely

func (e *Tonegeneratesrc) GetRepeat() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("repeat")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetRepeat(value bool) error {
	return e.Element.SetProperty("repeat", value)
}

// samplesperbuffer (int)
//
// Number of samples in each outgoing buffer

func (e *Tonegeneratesrc) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

// typefind (bool)
//
// Run typefind before negotiating (deprecated, non-functional)

func (e *Tonegeneratesrc) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

// volume (int)
//
// Volume of first signal

func (e *Tonegeneratesrc) GetVolume() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetVolume(value int) error {
	return e.Element.SetProperty("volume", value)
}

// volume2 (int)
//
// Volume of second tone signal

func (e *Tonegeneratesrc) GetVolume2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("volume2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Tonegeneratesrc) SetVolume2(value int) error {
	return e.Element.SetProperty("volume2", value)
}

