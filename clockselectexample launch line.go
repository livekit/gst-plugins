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

type ClockselectexampleLaunchLine struct {
	Element *gst.Element
}

func NewClockselectexampleLaunchLine() (*ClockselectexampleLaunchLine, error) {
	e, err := gst.NewElement("clockselectExample launch line")
	if err != nil {
		return nil, err
	}
	return &ClockselectexampleLaunchLine{Element: e}, nil
}

func NewClockselectexampleLaunchLineWithName(name string) (*ClockselectexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("clockselectExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &ClockselectexampleLaunchLine{Element: e}, nil
}

// ----- Properties -----

// clock-id (GstClockSelectClockId)
//
// ID of pipeline clock

func (e *ClockselectexampleLaunchLine) GetClockId() (GstClockSelectClockId, error) {
	var value GstClockSelectClockId
	var ok bool
	v, err := e.Element.GetProperty("clock-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstClockSelectClockId)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstClockSelectClockId")
	}
	return value, nil
}

func (e *ClockselectexampleLaunchLine) SetClockId(value GstClockSelectClockId) error {
	e.Element.SetArg("clock-id", string(value))
	return nil
}

// ptp-domain (uint)
//
// PTP clock domain (meaningful only when Clock ID is PTP)

func (e *ClockselectexampleLaunchLine) GetPtpDomain() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ptp-domain")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *ClockselectexampleLaunchLine) SetPtpDomain(value uint) error {
	return e.Element.SetProperty("ptp-domain", value)
}

// ----- Constants -----

type GstClockSelectClockId string

const (
	GstClockSelectClockIdDefault GstClockSelectClockId = "default" // default (0) – Default (elected from elements) pipeline clock
	GstClockSelectClockIdMonotonic GstClockSelectClockId = "monotonic" // monotonic (1) – System monotonic clock
	GstClockSelectClockIdRealtime GstClockSelectClockId = "realtime" // realtime (2) – System realtime clock
	GstClockSelectClockIdPtp GstClockSelectClockId = "ptp" // ptp (3) – PTP clock
	GstClockSelectClockIdTai GstClockSelectClockId = "tai" // tai (4) – System TAI clock
)

