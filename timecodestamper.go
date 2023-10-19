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

type Timecodestamper struct {
	*base.GstBaseTransform
}

func NewTimecodestamper() (*Timecodestamper, error) {
	e, err := gst.NewElement("timecodestamper")
	if err != nil {
		return nil, err
	}
	return &Timecodestamper{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewTimecodestamperWithName(name string) (*Timecodestamper, error) {
	e, err := gst.NewElementWithName("timecodestamper", name)
	if err != nil {
		return nil, err
	}
	return &Timecodestamper{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// auto-resync (bool)
//
// If true resync last known timecode from upstream, otherwise only count up from the last known one

func (e *Timecodestamper) GetAutoResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Timecodestamper) SetAutoResync(value bool) error {
	return e.Element.SetProperty("auto-resync", value)
}

// drop-frame (bool)
//
// Use drop-frame timecodes for 29.97 and 59.94 FPS

func (e *Timecodestamper) GetDropFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Timecodestamper) SetDropFrame(value bool) error {
	return e.Element.SetProperty("drop-frame", value)
}

// ltc-auto-resync (bool)
//
// If true the LTC timecode will be automatically resynced if it drifts, otherwise it will only be counted up from the last known one

func (e *Timecodestamper) GetLtcAutoResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ltc-auto-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Timecodestamper) SetLtcAutoResync(value bool) error {
	return e.Element.SetProperty("ltc-auto-resync", value)
}

// ltc-daily-jam (GstGdateTime)
//
// The daily jam of the LTC timecode

func (e *Timecodestamper) GetLtcDailyJam() (interface{}, error) {
	return e.Element.GetProperty("ltc-daily-jam")
}

func (e *Timecodestamper) SetLtcDailyJam(value interface{}) error {
	return e.Element.SetProperty("ltc-daily-jam", value)
}

// ltc-extra-latency (uint64)
//
// Extra latency to introduce for waiting for LTC timecodes

func (e *Timecodestamper) GetLtcExtraLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ltc-extra-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Timecodestamper) SetLtcExtraLatency(value uint64) error {
	return e.Element.SetProperty("ltc-extra-latency", value)
}

// ltc-timeout (uint64)
//
// Time out LTC timecode if no new timecode was detected after this time

func (e *Timecodestamper) GetLtcTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ltc-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Timecodestamper) SetLtcTimeout(value uint64) error {
	return e.Element.SetProperty("ltc-timeout", value)
}

// post-messages (bool)
//
// Post element message containing the current timecode

func (e *Timecodestamper) GetPostMessages() (bool, error) {
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

func (e *Timecodestamper) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

// rtc-auto-resync (bool)
//
// If true the RTC timecode will be automatically resynced if it drifts, otherwise it will only be counted up from the last known one

func (e *Timecodestamper) GetRtcAutoResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtc-auto-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Timecodestamper) SetRtcAutoResync(value bool) error {
	return e.Element.SetProperty("rtc-auto-resync", value)
}

// rtc-max-drift (uint64)
//
// Maximum number of nanoseconds the RTC clock is allowed to drift from the video before it is resynced

func (e *Timecodestamper) GetRtcMaxDrift() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("rtc-max-drift")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Timecodestamper) SetRtcMaxDrift(value uint64) error {
	return e.Element.SetProperty("rtc-max-drift", value)
}

// set (GstTimeCodeStamperSet)
//
// Choose whether timecodes should be overridden or not

func (e *Timecodestamper) GetSet() (GstTimeCodeStamperSet, error) {
	var value GstTimeCodeStamperSet
	var ok bool
	v, err := e.Element.GetProperty("set")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTimeCodeStamperSet)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTimeCodeStamperSet")
	}
	return value, nil
}

func (e *Timecodestamper) SetSet(value GstTimeCodeStamperSet) error {
	e.Element.SetArg("set", string(value))
	return nil
}

// set-internal-timecode (GstVideoTimeCode)
//
// If set, take this timecode as the internal timecode for the first frame and increment from it. Only the values itself and daily jam are taken, flags and frame rate are always determined by timecodestamper itself. If unset, the internal timecode will start at 0 with the daily jam being the current real-time clock time

func (e *Timecodestamper) GetSetInternalTimecode() (interface{}, error) {
	return e.Element.GetProperty("set-internal-timecode")
}

func (e *Timecodestamper) SetSetInternalTimecode(value interface{}) error {
	return e.Element.SetProperty("set-internal-timecode", value)
}

// source (GstTimeCodeStamperSource)
//
// Choose from what source the timecode should be taken

func (e *Timecodestamper) GetSource() (GstTimeCodeStamperSource, error) {
	var value GstTimeCodeStamperSource
	var ok bool
	v, err := e.Element.GetProperty("source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTimeCodeStamperSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTimeCodeStamperSource")
	}
	return value, nil
}

func (e *Timecodestamper) SetSource(value GstTimeCodeStamperSource) error {
	e.Element.SetArg("source", string(value))
	return nil
}

// timecode-offset (int)
//
// Add this offset in frames to internal, LTC or RTC timecode, useful if there is an offset between the timecode source and video

func (e *Timecodestamper) GetTimecodeOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timecode-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Timecodestamper) SetTimecodeOffset(value int) error {
	return e.Element.SetProperty("timecode-offset", value)
}

// timeout (uint64)
//
// Time out upstream timecode if no new timecode was detected after this time

func (e *Timecodestamper) GetTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Timecodestamper) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

// ----- Constants -----

type GstTimeCodeStamperSet string

const (
	GstTimeCodeStamperSetNever GstTimeCodeStamperSet = "never" // never (0) – Never set timecodes
	GstTimeCodeStamperSetKeep GstTimeCodeStamperSet = "keep" // keep (1) – Keep upstream timecodes and only set if no upstream timecode
	GstTimeCodeStamperSetAlways GstTimeCodeStamperSet = "always" // always (2) – Always set timecode and remove upstream timecode
)

type GstTimeCodeStamperSource string

const (
	GstTimeCodeStamperSourceInternal GstTimeCodeStamperSource = "internal" // internal (0) – Use internal timecode counter, starting at zero or value set by property
	GstTimeCodeStamperSourceZero GstTimeCodeStamperSource = "zero" // zero (1) – Always use zero
	GstTimeCodeStamperSourceLastKnown GstTimeCodeStamperSource = "last-known" // last-known (2) – Count up from the last known upstream timecode or internal if unknown
	GstTimeCodeStamperSourceLastKnownOrZero GstTimeCodeStamperSource = "last-known-or-zero" // last-known-or-zero (3) – Count up from the last known upstream timecode or zero if unknown
	GstTimeCodeStamperSourceLtc GstTimeCodeStamperSource = "ltc" // ltc (4) – Linear timecode from an audio device
	GstTimeCodeStamperSourceRtc GstTimeCodeStamperSource = "rtc" // rtc (5) – Timecode from real time clock
)

