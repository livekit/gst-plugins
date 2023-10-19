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

type Avwait struct {
	Element *gst.Element
}

func NewAvwait() (*Avwait, error) {
	e, err := gst.NewElement("avwait")
	if err != nil {
		return nil, err
	}
	return &Avwait{Element: e}, nil
}

func NewAvwaitWithName(name string) (*Avwait, error) {
	e, err := gst.NewElementWithName("avwait", name)
	if err != nil {
		return nil, err
	}
	return &Avwait{Element: e}, nil
}

// ----- Properties -----

// end-running-time (uint64)
//
// Running time to end at in running-time mode

func (e *Avwait) GetEndRunningTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("end-running-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Avwait) SetEndRunningTime(value uint64) error {
	return e.Element.SetProperty("end-running-time", value)
}

// end-timecode (GstVideoTimeCode)
//
// Timecode to end at in timecode mode (object)

func (e *Avwait) GetEndTimecode() (interface{}, error) {
	return e.Element.GetProperty("end-timecode")
}

func (e *Avwait) SetEndTimecode(value interface{}) error {
	return e.Element.SetProperty("end-timecode", value)
}

// mode (GstAvWaitMode)
//
// Operation mode: What to wait for

func (e *Avwait) GetMode() (GstAvWaitMode, error) {
	var value GstAvWaitMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvWaitMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvWaitMode")
	}
	return value, nil
}

func (e *Avwait) SetMode(value GstAvWaitMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// recording (bool)
//
// Whether the element is stopped or recording. If set to FALSE, all buffers will be dropped regardless of settings.

func (e *Avwait) GetRecording() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("recording")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Avwait) SetRecording(value bool) error {
	return e.Element.SetProperty("recording", value)
}

// target-running-time (uint64)
//
// Running time to wait for in running-time mode

func (e *Avwait) GetTargetRunningTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("target-running-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Avwait) SetTargetRunningTime(value uint64) error {
	return e.Element.SetProperty("target-running-time", value)
}

// target-timecode (GstVideoTimeCode)
//
// Timecode to wait for in timecode mode (object)

func (e *Avwait) GetTargetTimecode() (interface{}, error) {
	return e.Element.GetProperty("target-timecode")
}

func (e *Avwait) SetTargetTimecode(value interface{}) error {
	return e.Element.SetProperty("target-timecode", value)
}

// target-timecode-string (string)
//
// Timecode to wait for in timecode mode (string). Must take the form 00:00:00:00

func (e *Avwait) GetTargetTimecodeString() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("target-timecode-string")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Avwait) SetTargetTimecodeString(value string) error {
	return e.Element.SetProperty("target-timecode-string", value)
}

// ----- Constants -----

type GstAvWaitMode string

const (
	GstAvWaitModeTimecode GstAvWaitMode = "timecode" // timecode (0) – time code (default)
	GstAvWaitModeRunningTime GstAvWaitMode = "running-time" // running-time (1) – running time
	GstAvWaitModeVideoFirst GstAvWaitMode = "video-first" // video-first (2) – video first
)

