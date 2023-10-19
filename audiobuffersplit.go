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

type Audiobuffersplit struct {
	Element *gst.Element
}

func NewAudiobuffersplit() (*Audiobuffersplit, error) {
	e, err := gst.NewElement("audiobuffersplit")
	if err != nil {
		return nil, err
	}
	return &Audiobuffersplit{Element: e}, nil
}

func NewAudiobuffersplitWithName(name string) (*Audiobuffersplit, error) {
	e, err := gst.NewElementWithName("audiobuffersplit", name)
	if err != nil {
		return nil, err
	}
	return &Audiobuffersplit{Element: e}, nil
}

// ----- Properties -----

// alignment-threshold (uint64)
//
// Timestamp alignment threshold in nanoseconds

func (e *Audiobuffersplit) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audiobuffersplit) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

// discont-wait (uint64)
//
// Window of time in nanoseconds to wait before creating a discontinuity

func (e *Audiobuffersplit) GetDiscontWait() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("discont-wait")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audiobuffersplit) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

// gapless (bool)
//
// Insert silence/drop samples instead of creating a discontinuity

func (e *Audiobuffersplit) GetGapless() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("gapless")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audiobuffersplit) SetGapless(value bool) error {
	return e.Element.SetProperty("gapless", value)
}

// max-silence-time (uint64)
//
// Do not insert silence in gapless mode if the gap exceeds this period (in ns) (0 = disabled)

func (e *Audiobuffersplit) GetMaxSilenceTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-silence-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audiobuffersplit) SetMaxSilenceTime(value uint64) error {
	return e.Element.SetProperty("max-silence-time", value)
}

// output-buffer-duration (GstFraction)
//
// Output block size in seconds

func (e *Audiobuffersplit) GetOutputBufferDuration() (interface{}, error) {
	return e.Element.GetProperty("output-buffer-duration")
}

func (e *Audiobuffersplit) SetOutputBufferDuration(value interface{}) error {
	return e.Element.SetProperty("output-buffer-duration", value)
}

// output-buffer-size (uint)
//
// Allow specifying a buffer size for splitting. Zero by default.
// Takes precedence over output-buffer-duration when set to a
// non zero value else will not be in effect.

func (e *Audiobuffersplit) GetOutputBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("output-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Audiobuffersplit) SetOutputBufferSize(value uint) error {
	return e.Element.SetProperty("output-buffer-size", value)
}

// strict-buffer-size (bool)
//
// Discard the last samples at EOS or discont if they are too small to fill a buffer

func (e *Audiobuffersplit) GetStrictBufferSize() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("strict-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audiobuffersplit) SetStrictBufferSize(value bool) error {
	return e.Element.SetProperty("strict-buffer-size", value)
}

