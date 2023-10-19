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

type Glsinkbin struct {
	Element *gst.Element
}

func NewGlsinkbin() (*Glsinkbin, error) {
	e, err := gst.NewElement("glsinkbin")
	if err != nil {
		return nil, err
	}
	return &Glsinkbin{Element: e}, nil
}

func NewGlsinkbinWithName(name string) (*Glsinkbin, error) {
	e, err := gst.NewElementWithName("glsinkbin", name)
	if err != nil {
		return nil, err
	}
	return &Glsinkbin{Element: e}, nil
}

// ----- Properties -----

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Glsinkbin) GetAsync() (bool, error) {
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

func (e *Glsinkbin) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Glsinkbin) GetBlocksize() (uint, error) {
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

func (e *Glsinkbin) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// brightness (float64)
//
// brightness

func (e *Glsinkbin) GetBrightness() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glsinkbin) SetBrightness(value float64) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (float64)
//
// contrast

func (e *Glsinkbin) GetContrast() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glsinkbin) SetContrast(value float64) error {
	return e.Element.SetProperty("contrast", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Glsinkbin) GetEnableLastSample() (bool, error) {
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

func (e *Glsinkbin) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Glsinkbin) GetForceAspectRatio() (bool, error) {
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

func (e *Glsinkbin) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// hue (float64)
//
// hue

func (e *Glsinkbin) GetHue() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glsinkbin) SetHue(value float64) error {
	return e.Element.SetProperty("hue", value)
}

// last-sample (GstSample)
//
// The last sample received in the sink

func (e *Glsinkbin) GetLastSample() (*gst.Sample, error) {
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

func (e *Glsinkbin) GetMaxBitrate() (uint64, error) {
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

func (e *Glsinkbin) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Glsinkbin) GetMaxLateness() (int64, error) {
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

func (e *Glsinkbin) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Glsinkbin) GetQos() (bool, error) {
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

func (e *Glsinkbin) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Glsinkbin) GetRenderDelay() (uint64, error) {
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

func (e *Glsinkbin) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// saturation (float64)
//
// saturation

func (e *Glsinkbin) GetSaturation() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Glsinkbin) SetSaturation(value float64) error {
	return e.Element.SetProperty("saturation", value)
}

// sink (GstElement)
//
// The GL sink chain to use

func (e *Glsinkbin) GetSink() (interface{}, error) {
	return e.Element.GetProperty("sink")
}

func (e *Glsinkbin) SetSink(value interface{}) error {
	return e.Element.SetProperty("sink", value)
}

// sync (bool)
//
// Sync on the clock

func (e *Glsinkbin) GetSync() (bool, error) {
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

func (e *Glsinkbin) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Glsinkbin) GetThrottleTime() (uint64, error) {
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

func (e *Glsinkbin) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Glsinkbin) GetTsOffset() (int64, error) {
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

func (e *Glsinkbin) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

