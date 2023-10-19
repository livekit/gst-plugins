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

type Tinyalsasink struct {
	*base.GstBaseSink
}

func NewTinyalsasink() (*Tinyalsasink, error) {
	e, err := gst.NewElement("tinyalsasink")
	if err != nil {
		return nil, err
	}
	return &Tinyalsasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewTinyalsasinkWithName(name string) (*Tinyalsasink, error) {
	e, err := gst.NewElementWithName("tinyalsasink", name)
	if err != nil {
		return nil, err
	}
	return &Tinyalsasink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// alignment-threshold (uint64)
//
// Timestamp alignment threshold in nanoseconds

func (e *Tinyalsasink) GetAlignmentThreshold() (uint64, error) {
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

func (e *Tinyalsasink) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Tinyalsasink) GetAsync() (bool, error) {
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

func (e *Tinyalsasink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Tinyalsasink) GetBlocksize() (uint, error) {
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

func (e *Tinyalsasink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// buffer-time (int64)
//
// Size of audio buffer in microseconds, this is the minimum latency that the sink reports

func (e *Tinyalsasink) GetBufferTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Tinyalsasink) SetBufferTime(value int64) error {
	return e.Element.SetProperty("buffer-time", value)
}

// can-activate-pull (bool)
//
// Allow pull-based scheduling

func (e *Tinyalsasink) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tinyalsasink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// card (uint)
//
// The ALSA card to use

func (e *Tinyalsasink) GetCard() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("card")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Tinyalsasink) SetCard(value uint) error {
	return e.Element.SetProperty("card", value)
}

// device (uint)
//
// The ALSA device to use

func (e *Tinyalsasink) GetDevice() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Tinyalsasink) SetDevice(value uint) error {
	return e.Element.SetProperty("device", value)
}

// discont-wait (uint64)
//
// Window of time in nanoseconds to wait before creating a discontinuity

func (e *Tinyalsasink) GetDiscontWait() (uint64, error) {
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

func (e *Tinyalsasink) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

// drift-tolerance (int64)
//
// Tolerance for clock drift in microseconds

func (e *Tinyalsasink) GetDriftTolerance() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("drift-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Tinyalsasink) SetDriftTolerance(value int64) error {
	return e.Element.SetProperty("drift-tolerance", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Tinyalsasink) GetEnableLastSample() (bool, error) {
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

func (e *Tinyalsasink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// last-sample (GstSample)
//
// The last sample received in the sink

func (e *Tinyalsasink) GetLastSample() (*gst.Sample, error) {
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

// latency-time (int64)
//
// The minimum amount of data to write in each iteration in microseconds

func (e *Tinyalsasink) GetLatencyTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("latency-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Tinyalsasink) SetLatencyTime(value int64) error {
	return e.Element.SetProperty("latency-time", value)
}

// max-bitrate (uint64)
//
// The maximum bits per second to render (0 = disabled)

func (e *Tinyalsasink) GetMaxBitrate() (uint64, error) {
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

func (e *Tinyalsasink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Tinyalsasink) GetMaxLateness() (int64, error) {
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

func (e *Tinyalsasink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Tinyalsasink) GetProcessingDeadline() (uint64, error) {
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

func (e *Tinyalsasink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// provide-clock (bool)
//
// Provide a clock to be used as the global pipeline clock

func (e *Tinyalsasink) GetProvideClock() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("provide-clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Tinyalsasink) SetProvideClock(value bool) error {
	return e.Element.SetProperty("provide-clock", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Tinyalsasink) GetQos() (bool, error) {
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

func (e *Tinyalsasink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Tinyalsasink) GetRenderDelay() (uint64, error) {
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

func (e *Tinyalsasink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// slave-method (GstAudioBaseSinkSlaveMethod)
//
// Algorithm used to match the rate of the masterclock

func (e *Tinyalsasink) GetSlaveMethod() (interface{}, error) {
	return e.Element.GetProperty("slave-method")
}

func (e *Tinyalsasink) SetSlaveMethod(value interface{}) error {
	return e.Element.SetProperty("slave-method", value)
}

// stats (GstStructure)
//
// Sink Statistics

func (e *Tinyalsasink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// sync (bool)
//
// Sync on the clock

func (e *Tinyalsasink) GetSync() (bool, error) {
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

func (e *Tinyalsasink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Tinyalsasink) GetThrottleTime() (uint64, error) {
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

func (e *Tinyalsasink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Tinyalsasink) GetTsOffset() (int64, error) {
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

func (e *Tinyalsasink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

