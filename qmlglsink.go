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

type Qmlglsink struct {
	*base.GstBaseSink
}

func NewQmlglsink() (*Qmlglsink, error) {
	e, err := gst.NewElement("qmlglsink")
	if err != nil {
		return nil, err
	}
	return &Qmlglsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewQmlglsinkWithName(name string) (*Qmlglsink, error) {
	e, err := gst.NewElementWithName("qmlglsink", name)
	if err != nil {
		return nil, err
	}
	return &Qmlglsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Qmlglsink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Qmlglsink) GetBlocksize() (uint, error) {
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

func (e *Qmlglsink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Qmlglsink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Qmlglsink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// last-sample (GstSample)
//
// The last sample received in the sink

func (e *Qmlglsink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// max-bitrate (uint64)
//
// The maximum bits per second to render (0 = disabled)

func (e *Qmlglsink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Qmlglsink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Qmlglsink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Qmlglsink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// pixel-aspect-ratio (GstFraction)
//
// The pixel aspect ratio of the device

func (e *Qmlglsink) GetPixelAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("pixel-aspect-ratio")
}

func (e *Qmlglsink) SetPixelAspectRatio(value interface{}) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Qmlglsink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Qmlglsink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Qmlglsink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Qmlglsink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Qmlglsink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// show-preroll-frame (bool)
//
// Whether to render video frames during preroll

func (e *Qmlglsink) GetShowPrerollFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-preroll-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsink) SetShowPrerollFrame(value bool) error {
	return e.Element.SetProperty("show-preroll-frame", value)
}

// stats (GstStructure)
//
// Sink Statistics

func (e *Qmlglsink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// sync (bool)
//
// Sync on the clock

func (e *Qmlglsink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Qmlglsink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Qmlglsink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Qmlglsink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Qmlglsink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

// widget (GstGpointer)
//
// The QQuickItem to place in the object hierarchy

func (e *Qmlglsink) GetWidget() (interface{}, error) {
	return e.Element.GetProperty("widget")
}

func (e *Qmlglsink) SetWidget(value interface{}) error {
	return e.Element.SetProperty("widget", value)
}

