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

type Queue struct {
	Element *gst.Element
}

func NewQueue() (*Queue, error) {
	e, err := gst.NewElement("queue")
	if err != nil {
		return nil, err
	}
	return &Queue{Element: e}, nil
}

func NewQueueWithName(name string) (*Queue, error) {
	e, err := gst.NewElementWithName("queue", name)
	if err != nil {
		return nil, err
	}
	return &Queue{Element: e}, nil
}

// ----- Properties -----

// current-level-buffers (uint)
//
// Current number of buffers in the queue

func (e *Queue) GetCurrentLevelBuffers() (uint, error) {
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

func (e *Queue) GetCurrentLevelBytes() (uint, error) {
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

func (e *Queue) GetCurrentLevelTime() (uint64, error) {
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

// flush-on-eos (bool)
//
// Discard all data in the queue when an EOS event is received, and pass
// on the EOS event as soon as possible (instead of waiting until all
// buffers in the queue have been processed, which is the default behaviour).

// leaky (GstQueueLeaky)
//
// Where the queue leaks, if at all

func (e *Queue) GetLeaky() (GstQueueLeaky, error) {
	var value GstQueueLeaky
	var ok bool
	v, err := e.Element.GetProperty("leaky")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQueueLeaky)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQueueLeaky")
	}
	return value, nil
}

func (e *Queue) SetLeaky(value GstQueueLeaky) error {
	e.Element.SetArg("leaky", string(value))
	return nil
}

// max-size-buffers (uint)
//
// Max. number of buffers in the queue (0=disable)

func (e *Queue) GetMaxSizeBuffers() (uint, error) {
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

func (e *Queue) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

// max-size-bytes (uint)
//
// Max. amount of data in the queue (bytes, 0=disable)

func (e *Queue) GetMaxSizeBytes() (uint, error) {
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

func (e *Queue) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

// max-size-time (uint64)
//
// Max. amount of data in the queue (in ns, 0=disable)

func (e *Queue) GetMaxSizeTime() (uint64, error) {
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

func (e *Queue) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

// min-threshold-buffers (uint)
//
// Min. number of buffers in the queue to allow reading (0=disable)

func (e *Queue) GetMinThresholdBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-threshold-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Queue) SetMinThresholdBuffers(value uint) error {
	return e.Element.SetProperty("min-threshold-buffers", value)
}

// min-threshold-bytes (uint)
//
// Min. amount of data in the queue to allow reading (bytes, 0=disable)

func (e *Queue) GetMinThresholdBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-threshold-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Queue) SetMinThresholdBytes(value uint) error {
	return e.Element.SetProperty("min-threshold-bytes", value)
}

// min-threshold-time (uint64)
//
// Min. amount of data in the queue to allow reading (in ns, 0=disable)

func (e *Queue) GetMinThresholdTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-threshold-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Queue) SetMinThresholdTime(value uint64) error {
	return e.Element.SetProperty("min-threshold-time", value)
}

// silent (bool)
//
// Don't emit queue signals. Makes queues more lightweight if no signals are
// needed.

func (e *Queue) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Queue) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// ----- Constants -----

type GstQueueLeaky string

const (
	GstQueueLeakyNo GstQueueLeaky = "no" // no (0) – Not Leaky
	GstQueueLeakyUpstream GstQueueLeaky = "upstream" // upstream (1) – Leaky on upstream (new buffers)
	GstQueueLeakyDownstream GstQueueLeaky = "downstream" // downstream (2) – Leaky on downstream (old buffers)
)

