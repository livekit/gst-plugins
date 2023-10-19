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

package base

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type GstBaseSink struct {
	Element *gst.Element
}

// ----- Properties -----

// async (bool)
//
// If set to TRUE, the basesink will perform asynchronous state changes.
// When set to FALSE, the sink will not signal the parent when it prerolls.
// Use this option when dealing with sparse streams or when synchronisation is
// not required.

func (e *GstBaseSink) GetAsync() (bool, error) {
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

func (e *GstBaseSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// The amount of bytes to pull when operating in pull mode.

func (e *GstBaseSink) GetBlocksize() (uint, error) {
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

func (e *GstBaseSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property. If FALSE, basesink doesn't keep a
// reference to the last buffer arrived and the last-sample property is always
// set to NULL. This can be useful if you need buffers to be released as soon
// as possible, eg. if you're using a buffer pool.

func (e *GstBaseSink) GetEnableLastSample() (bool, error) {
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

func (e *GstBaseSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// last-sample (GstSample)
//
// The last buffer that arrived in the sink and was used for preroll or for
// rendering. This property can be used to generate thumbnails. This property
// can be NULL when the sink has not yet received a buffer.

func (e *GstBaseSink) GetLastSample() (*gst.Sample, error) {
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
// Control the maximum amount of bits that will be rendered per second.
// Setting this property to a value bigger than 0 will make the sink delay
// rendering of the buffers when it would exceed to max-bitrate.

func (e *GstBaseSink) GetMaxBitrate() (uint64, error) {
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

func (e *GstBaseSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)

func (e *GstBaseSink) GetMaxLateness() (int64, error) {
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

func (e *GstBaseSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// processing-deadline (uint64)
//
// Maximum amount of time (in nanoseconds) that the pipeline can take
// for processing the buffer. This is added to the latency of live
// pipelines.

func (e *GstBaseSink) GetProcessingDeadline() (uint64, error) {
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

func (e *GstBaseSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// qos (bool)

func (e *GstBaseSink) GetQos() (bool, error) {
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

func (e *GstBaseSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// The additional delay between synchronisation and actual rendering of the
// media. This property will add additional latency to the device in order to
// make other sinks compensate for the delay.

func (e *GstBaseSink) GetRenderDelay() (uint64, error) {
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

func (e *GstBaseSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// stats (GstStructure)
//
// Various GstBaseSink statistics. This property returns a GstStructure
// with name application/x-gst-base-sink-stats with the following fields:

// sync (bool)

func (e *GstBaseSink) GetSync() (bool, error) {
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

func (e *GstBaseSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to insert between buffers. This property can be used to control
// the maximum amount of buffers per second to render. Setting this property
// to a value bigger than 0 will make the sink create THROTTLE QoS events.

func (e *GstBaseSink) GetThrottleTime() (uint64, error) {
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

func (e *GstBaseSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Controls the final synchronisation, a negative value will render the buffer
// earlier while a positive value delays playback. This property can be
// used to fix synchronisation in bad files.

func (e *GstBaseSink) GetTsOffset() (int64, error) {
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

func (e *GstBaseSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

