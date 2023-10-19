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

type Videorate struct {
	*base.GstBaseTransform
}

func NewVideorate() (*Videorate, error) {
	e, err := gst.NewElement("videorate")
	if err != nil {
		return nil, err
	}
	return &Videorate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVideorateWithName(name string) (*Videorate, error) {
	e, err := gst.NewElementWithName("videorate", name)
	if err != nil {
		return nil, err
	}
	return &Videorate{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// average-period (uint64)
//
// Arrange for maximum framerate by dropping frames beyond a certain framerate,
// where the framerate is calculated using a moving average over the
// configured.

func (e *Videorate) GetAveragePeriod() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("average-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Videorate) SetAveragePeriod(value uint64) error {
	return e.Element.SetProperty("average-period", value)
}

// drop (uint64)
//
// Number of dropped frames

func (e *Videorate) GetDrop() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// drop-only (bool)
//
// Only drop frames, no duplicates are produced.

func (e *Videorate) GetDropOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videorate) SetDropOnly(value bool) error {
	return e.Element.SetProperty("drop-only", value)
}

// duplicate (uint64)
//
// Number of duplicated frames

func (e *Videorate) GetDuplicate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("duplicate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// in (uint64)
//
// Number of input frames

func (e *Videorate) GetIn() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("in")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// max-closing-segment-duplication-duration (uint64)
//
// Limits the maximum duration for which the last buffer is duplicated when
// finalizing a segment or on EOS. When receiving an EOS event or a new
// segment, videorate duplicates the last frame to close the configured
// segment (copying the last buffer until its stop time (or
// start time for reverse playback) is reached), this property
// ensures that it won't push buffers covering a duration longer than
// specified.

func (e *Videorate) GetMaxClosingSegmentDuplicationDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-closing-segment-duplication-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Videorate) SetMaxClosingSegmentDuplicationDuration(value uint64) error {
	return e.Element.SetProperty("max-closing-segment-duplication-duration", value)
}

// max-duplication-time (uint64)
//
// Duplicate frames only if the gap between two consecutive frames does not
// exceed this duration.

func (e *Videorate) GetMaxDuplicationTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-duplication-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Videorate) SetMaxDuplicationTime(value uint64) error {
	return e.Element.SetProperty("max-duplication-time", value)
}

// max-rate (int)
//
// maximum framerate to pass through

func (e *Videorate) GetMaxRate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Videorate) SetMaxRate(value int) error {
	return e.Element.SetProperty("max-rate", value)
}

// new-pref (float64)
//
// Value indicating how much to prefer new frames (unused)

func (e *Videorate) GetNewPref() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("new-pref")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videorate) SetNewPref(value float64) error {
	return e.Element.SetProperty("new-pref", value)
}

// out (uint64)
//
// Number of output frames

func (e *Videorate) GetOut() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("out")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// rate (float64)
//
// Factor of speed for frame displaying

func (e *Videorate) GetRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Videorate) SetRate(value float64) error {
	return e.Element.SetProperty("rate", value)
}

// silent (bool)
//
// Don't emit notify for dropped and duplicated frames

func (e *Videorate) GetSilent() (bool, error) {
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

func (e *Videorate) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// skip-to-first (bool)
//
// Don't produce buffers before the first one we receive.

func (e *Videorate) GetSkipToFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-to-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Videorate) SetSkipToFirst(value bool) error {
	return e.Element.SetProperty("skip-to-first", value)
}

