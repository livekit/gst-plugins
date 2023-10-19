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

type Dashsink struct {
	Element *gst.Element
}

func NewDashsink() (*Dashsink, error) {
	e, err := gst.NewElement("dashsink")
	if err != nil {
		return nil, err
	}
	return &Dashsink{Element: e}, nil
}

func NewDashsinkWithName(name string) (*Dashsink, error) {
	e, err := gst.NewElementWithName("dashsink", name)
	if err != nil {
		return nil, err
	}
	return &Dashsink{Element: e}, nil
}

// ----- Properties -----

// dynamic (bool)
//
// Provides a dynamic mpd

func (e *Dashsink) GetDynamic() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dynamic")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dashsink) SetDynamic(value bool) error {
	return e.Element.SetProperty("dynamic", value)
}

// min-buffer-time (uint64)
//
// Provides to the manifest a minimum buffer time in milliseconds

func (e *Dashsink) GetMinBufferTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dashsink) SetMinBufferTime(value uint64) error {
	return e.Element.SetProperty("min-buffer-time", value)
}

// minimum-update-period (uint64)
//
// Provides to the manifest a minimum update period in milliseconds

func (e *Dashsink) GetMinimumUpdatePeriod() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("minimum-update-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dashsink) SetMinimumUpdatePeriod(value uint64) error {
	return e.Element.SetProperty("minimum-update-period", value)
}

// mpd-baseurl (string)
//
// BaseURL to set in the MPD

func (e *Dashsink) GetMpdBaseurl() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mpd-baseurl")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dashsink) SetMpdBaseurl(value string) error {
	return e.Element.SetProperty("mpd-baseurl", value)
}

// mpd-filename (string)
//
// filename of the mpd to write

func (e *Dashsink) GetMpdFilename() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mpd-filename")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dashsink) SetMpdFilename(value string) error {
	return e.Element.SetProperty("mpd-filename", value)
}

// mpd-root-path (string)
//
// Path where the MPD and its fragents will be written

func (e *Dashsink) GetMpdRootPath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mpd-root-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dashsink) SetMpdRootPath(value string) error {
	return e.Element.SetProperty("mpd-root-path", value)
}

// muxer (GstDashSinkMuxerType)
//
// Muxer type to be used by dashsink to generate the fragment

func (e *Dashsink) GetMuxer() (GstDashSinkMuxerType, error) {
	var value GstDashSinkMuxerType
	var ok bool
	v, err := e.Element.GetProperty("muxer")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDashSinkMuxerType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDashSinkMuxerType")
	}
	return value, nil
}

func (e *Dashsink) SetMuxer(value GstDashSinkMuxerType) error {
	e.Element.SetArg("muxer", string(value))
	return nil
}

// period-duration (uint64)
//
// Provides the explicit duration of a period in milliseconds

func (e *Dashsink) GetPeriodDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("period-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dashsink) SetPeriodDuration(value uint64) error {
	return e.Element.SetProperty("period-duration", value)
}

// send-keyframe-requests (bool)
//
// Send keyframe requests to ensure correct fragmentation. If this is disabled then the input must have keyframes in regular intervals

func (e *Dashsink) GetSendKeyframeRequests() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-keyframe-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dashsink) SetSendKeyframeRequests(value bool) error {
	return e.Element.SetProperty("send-keyframe-requests", value)
}

// suggested-presentation-delay (uint64)
//
// set suggested presentation delay of MPD file in milliseconds

func (e *Dashsink) GetSuggestedPresentationDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("suggested-presentation-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dashsink) SetSuggestedPresentationDelay(value uint64) error {
	return e.Element.SetProperty("suggested-presentation-delay", value)
}

// target-duration (uint)
//
// The target duration in seconds of a segment/file. (0 - disabled, useful for management of segment duration by the streaming server)

func (e *Dashsink) GetTargetDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("target-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dashsink) SetTargetDuration(value uint) error {
	return e.Element.SetProperty("target-duration", value)
}

// use-segment-list (bool)
//
// Use segment list instead of segment template to create the segments

func (e *Dashsink) GetUseSegmentList() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-segment-list")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dashsink) SetUseSegmentList(value bool) error {
	return e.Element.SetProperty("use-segment-list", value)
}

// ----- Constants -----

type GstDashSinkMuxerType string

const (
	GstDashSinkMuxerTypeTs GstDashSinkMuxerType = "ts" // ts (0) – Use mpegtsmux
	GstDashSinkMuxerTypeMp4 GstDashSinkMuxerType = "mp4" // mp4 (1) – Use mp4mux (deprecated, non-functional)
	GstDashSinkMuxerTypeDashmp4 GstDashSinkMuxerType = "dashmp4" // dashmp4 (2) – Use dashmp4mux
)

