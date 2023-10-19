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

type Appsrc struct {
	*base.GstBaseSrc
}

func NewAppsrc() (*Appsrc, error) {
	e, err := gst.NewElement("appsrc")
	if err != nil {
		return nil, err
	}
	return &Appsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewAppsrcWithName(name string) (*Appsrc, error) {
	e, err := gst.NewElementWithName("appsrc", name)
	if err != nil {
		return nil, err
	}
	return &Appsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// block (bool)
//
// When max-bytes are queued and after the enough-data signal has been emitted,
// block any further push-buffer calls until the amount of queued bytes drops
// below the max-bytes limit.

func (e *Appsrc) GetBlock() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("block")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Appsrc) SetBlock(value bool) error {
	return e.Element.SetProperty("block", value)
}

// caps (GstCaps)
//
// The GstCaps that will negotiated downstream and will be put
// on outgoing buffers.

func (e *Appsrc) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Appsrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// current-level-buffers (uint64)
//
// The number of currently queued buffers inside appsrc.

func (e *Appsrc) GetCurrentLevelBuffers() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// current-level-bytes (uint64)
//
// The number of currently queued bytes inside appsrc.

func (e *Appsrc) GetCurrentLevelBytes() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// current-level-time (uint64)
//
// The amount of currently queued time inside appsrc.

func (e *Appsrc) GetCurrentLevelTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// duration (uint64)
//
// The total duration in nanoseconds of the data stream. If the total duration is known, it
// is recommended to configure it with this property.

func (e *Appsrc) GetDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Appsrc) SetDuration(value uint64) error {
	return e.Element.SetProperty("duration", value)
}

// emit-signals (bool)
//
// Make appsrc emit the "need-data", "enough-data" and "seek-data" signals.
// This option is by default enabled for backwards compatibility reasons but
// can disabled when needed because signal emission is expensive.

func (e *Appsrc) GetEmitSignals() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("emit-signals")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Appsrc) SetEmitSignals(value bool) error {
	return e.Element.SetProperty("emit-signals", value)
}

// format (GstFormat)
//
// The format to use for segment events. When the source is producing
// timestamped buffers this property should be set to GST_FORMAT_TIME.

func (e *Appsrc) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

func (e *Appsrc) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

// handle-segment-change (bool)
//
// When enabled, appsrc will check GstSegment in GstSample which was
// pushed via gst_app_src_push_sample or "push-sample" signal action.
// If a GstSegment is changed, corresponding segment event will be followed
// by next data flow.

// is-live (bool)
//
// Instruct the source to behave like a live source. This includes that it
// will only push out buffers in the PLAYING state.

func (e *Appsrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Appsrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// leaky-type (GstAppLeakyType)
//
// When set to any other value than GST_APP_LEAKY_TYPE_NONE then the appsrc
// will drop any buffers that are pushed into it once its internal queue is
// full. The selected type defines whether to drop the oldest or new
// buffers.

func (e *Appsrc) GetLeakyType() (interface{}, error) {
	return e.Element.GetProperty("leaky-type")
}

func (e *Appsrc) SetLeakyType(value interface{}) error {
	return e.Element.SetProperty("leaky-type", value)
}

// max-buffers (uint64)
//
// The maximum amount of buffers that can be queued internally.
// After the maximum amount of buffers are queued, appsrc will emit the
// "enough-data" signal.

func (e *Appsrc) GetMaxBuffers() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Appsrc) SetMaxBuffers(value uint64) error {
	return e.Element.SetProperty("max-buffers", value)
}

// max-bytes (uint64)
//
// The maximum amount of bytes that can be queued internally.
// After the maximum amount of bytes are queued, appsrc will emit the
// "enough-data" signal.

func (e *Appsrc) GetMaxBytes() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Appsrc) SetMaxBytes(value uint64) error {
	return e.Element.SetProperty("max-bytes", value)
}

// max-latency (int64)
//
// The maximum latency (-1 = unlimited)

func (e *Appsrc) GetMaxLatency() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Appsrc) SetMaxLatency(value int64) error {
	return e.Element.SetProperty("max-latency", value)
}

// max-time (uint64)
//
// The maximum amount of time that can be queued internally.
// After the maximum amount of time are queued, appsrc will emit the
// "enough-data" signal.

func (e *Appsrc) GetMaxTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Appsrc) SetMaxTime(value uint64) error {
	return e.Element.SetProperty("max-time", value)
}

// min-latency (int64)
//
// The minimum latency of the source. A value of -1 will use the default
// latency calculations of GstBaseSrc.

func (e *Appsrc) GetMinLatency() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("min-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Appsrc) SetMinLatency(value int64) error {
	return e.Element.SetProperty("min-latency", value)
}

// min-percent (uint)
//
// Make appsrc emit the "need-data" signal when the amount of bytes in the
// queue drops below this percentage of max-bytes.

func (e *Appsrc) GetMinPercent() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Appsrc) SetMinPercent(value uint) error {
	return e.Element.SetProperty("min-percent", value)
}

// size (int64)
//
// The total size in bytes of the data stream. If the total size is known, it
// is recommended to configure it with this property.

func (e *Appsrc) GetSize() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Appsrc) SetSize(value int64) error {
	return e.Element.SetProperty("size", value)
}

// stream-type (GstAppStreamType)
//
// The type of stream that this source is producing.  For seekable streams the
// application should connect to the seek-data signal.

func (e *Appsrc) GetStreamType() (interface{}, error) {
	return e.Element.GetProperty("stream-type")
}

func (e *Appsrc) SetStreamType(value interface{}) error {
	return e.Element.SetProperty("stream-type", value)
}

