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

type Queue2 struct {
	Element *gst.Element
}

func NewQueue2() (*Queue2, error) {
	e, err := gst.NewElement("queue2")
	if err != nil {
		return nil, err
	}
	return &Queue2{Element: e}, nil
}

func NewQueue2WithName(name string) (*Queue2, error) {
	e, err := gst.NewElementWithName("queue2", name)
	if err != nil {
		return nil, err
	}
	return &Queue2{Element: e}, nil
}

// ----- Properties -----

// avg-in-rate (int64)
//
// The average input data rate.

func (e *Queue2) GetAvgInRate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("avg-in-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// bitrate (uint64)
//
// The value used to convert between byte and time values for limiting
// the size of the queue.  Values are taken from either the upstream tags
// or from the downstream bitrate query.

func (e *Queue2) GetBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// current-level-buffers (uint)
//
// Current number of buffers in the queue

func (e *Queue2) GetCurrentLevelBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// current-level-bytes (uint)
//
// Current amount of data in the queue (bytes)

func (e *Queue2) GetCurrentLevelBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// current-level-time (uint64)
//
// Current amount of data in the queue (in ns)

func (e *Queue2) GetCurrentLevelTime() (uint64, error) {
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

// high-percent (int)
//
// High threshold for buffering to finish. Only used if use-buffering is True (Deprecated: use high-watermark instead)

func (e *Queue2) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Queue2) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

// high-watermark (float64)
//
// High threshold for buffering to finish. Only used if use-buffering is True

func (e *Queue2) GetHighWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Queue2) SetHighWatermark(value float64) error {
	return e.Element.SetProperty("high-watermark", value)
}

// low-percent (int)
//
// Low threshold for buffering to start. Only used if use-buffering is True (Deprecated: use low-watermark instead)

func (e *Queue2) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Queue2) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

// low-watermark (float64)
//
// Low threshold for buffering to start. Only used if use-buffering is True

func (e *Queue2) GetLowWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Queue2) SetLowWatermark(value float64) error {
	return e.Element.SetProperty("low-watermark", value)
}

// max-size-buffers (uint)
//
// Max. number of buffers in the queue (0=disable)

func (e *Queue2) GetMaxSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Queue2) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

// max-size-bytes (uint)
//
// Max. amount of data in the queue (bytes, 0=disable)

func (e *Queue2) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Queue2) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

// max-size-time (uint64)
//
// Max. amount of data in the queue (in ns, 0=disable)

func (e *Queue2) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Queue2) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

// ring-buffer-max-size (uint64)
//
// The maximum size of the ring buffer in bytes. If set to 0, the ring
// buffer is disabled. Default 0.

func (e *Queue2) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Queue2) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

// temp-location (string)
//
// Location to store temporary files in (Only read this property, use temp-template to configure the name template)

func (e *Queue2) GetTempLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// temp-remove (bool)
//
// When temp-template is set, remove the temporary file when going to READY.

func (e *Queue2) GetTempRemove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temp-remove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Queue2) SetTempRemove(value bool) error {
	return e.Element.SetProperty("temp-remove", value)
}

// temp-template (string)
//
// File template to store temporary files in, should contain directory and XXXXXX. (NULL == disabled)

func (e *Queue2) GetTempTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Queue2) SetTempTemplate(value string) error {
	return e.Element.SetProperty("temp-template", value)
}

// use-bitrate-query (bool)
//
// Use a bitrate from a downstream query to estimate buffer duration if not provided

func (e *Queue2) GetUseBitrateQuery() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-bitrate-query")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Queue2) SetUseBitrateQuery(value bool) error {
	return e.Element.SetProperty("use-bitrate-query", value)
}

// use-buffering (bool)
//
// Emit GST_MESSAGE_BUFFERING based on low-/high-percent thresholds (0%% = low-watermark, 100%% = high-watermark)

func (e *Queue2) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Queue2) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

// use-rate-estimate (bool)
//
// Estimate the bitrate of the stream to calculate time level

func (e *Queue2) GetUseRateEstimate() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-rate-estimate")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Queue2) SetUseRateEstimate(value bool) error {
	return e.Element.SetProperty("use-rate-estimate", value)
}

// use-tags-bitrate (bool)
//
// Use a bitrate from upstream tags to estimate buffer duration if not provided

func (e *Queue2) GetUseTagsBitrate() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-tags-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Queue2) SetUseTagsBitrate(value bool) error {
	return e.Element.SetProperty("use-tags-bitrate", value)
}

