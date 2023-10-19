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

type Dwritetimeoverlay struct {
	*base.GstBaseTransform
}

func NewDwritetimeoverlay() (*Dwritetimeoverlay, error) {
	e, err := gst.NewElement("dwritetimeoverlay")
	if err != nil {
		return nil, err
	}
	return &Dwritetimeoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDwritetimeoverlayWithName(name string) (*Dwritetimeoverlay, error) {
	e, err := gst.NewElementWithName("dwritetimeoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Dwritetimeoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// datetime-epoch (GstGdateTime)
//
// When showing times as dates, the initial date from which time is counted, if not specified prime epoch is used (1900-01-01)

func (e *Dwritetimeoverlay) GetDatetimeEpoch() (interface{}, error) {
	return e.Element.GetProperty("datetime-epoch")
}

func (e *Dwritetimeoverlay) SetDatetimeEpoch(value interface{}) error {
	return e.Element.SetProperty("datetime-epoch", value)
}

// datetime-format (string)
//
// When showing times as dates, the format to render date and time in

func (e *Dwritetimeoverlay) GetDatetimeFormat() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("datetime-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dwritetimeoverlay) SetDatetimeFormat(value string) error {
	return e.Element.SetProperty("datetime-format", value)
}

// reference-timestamp-caps (GstCaps)
//
// Caps to use for the reference timestamp time mode

func (e *Dwritetimeoverlay) GetReferenceTimestampCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("reference-timestamp-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Dwritetimeoverlay) SetReferenceTimestampCaps(value *gst.Caps) error {
	return e.Element.SetProperty("reference-timestamp-caps", value)
}

// show-times-as-dates (bool)
//
// Whether to display times, counted from datetime-epoch, as dates

func (e *Dwritetimeoverlay) GetShowTimesAsDates() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-times-as-dates")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritetimeoverlay) SetShowTimesAsDates(value bool) error {
	return e.Element.SetProperty("show-times-as-dates", value)
}

// time-mode (GstDwriteTimeOverlayTimeLine)
//
// What time to show

func (e *Dwritetimeoverlay) GetTimeMode() (GstDwriteTimeOverlayTimeLine, error) {
	var value GstDwriteTimeOverlayTimeLine
	var ok bool
	v, err := e.Element.GetProperty("time-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDwriteTimeOverlayTimeLine)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDwriteTimeOverlayTimeLine")
	}
	return value, nil
}

func (e *Dwritetimeoverlay) SetTimeMode(value GstDwriteTimeOverlayTimeLine) error {
	e.Element.SetArg("time-mode", string(value))
	return nil
}

// ----- Constants -----

type GstDwriteTimeOverlayTimeLine string

const (
	GstDwriteTimeOverlayTimeLineBufferTime GstDwriteTimeOverlayTimeLine = "buffer-time" // buffer-time (0) – buffer-time
	GstDwriteTimeOverlayTimeLineStreamTime GstDwriteTimeOverlayTimeLine = "stream-time" // stream-time (1) – stream-time
	GstDwriteTimeOverlayTimeLineRunningTime GstDwriteTimeOverlayTimeLine = "running-time" // running-time (2) – running-time
	GstDwriteTimeOverlayTimeLineTimeCode GstDwriteTimeOverlayTimeLine = "time-code" // time-code (3) – time-code
	GstDwriteTimeOverlayTimeLineElapsedRunningTime GstDwriteTimeOverlayTimeLine = "elapsed-running-time" // elapsed-running-time (4) – elapsed-running-time
	GstDwriteTimeOverlayTimeLineReferenceTimestamp GstDwriteTimeOverlayTimeLine = "reference-timestamp" // reference-timestamp (5) – reference-timestamp
	GstDwriteTimeOverlayTimeLineBufferCount GstDwriteTimeOverlayTimeLine = "buffer-count" // buffer-count (6) – buffer-count
	GstDwriteTimeOverlayTimeLineBufferOffset GstDwriteTimeOverlayTimeLine = "buffer-offset" // buffer-offset (7) – buffer-offset
)

