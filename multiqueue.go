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

type Multiqueue struct {
	Element *gst.Element
}

func NewMultiqueue() (*Multiqueue, error) {
	e, err := gst.NewElement("multiqueue")
	if err != nil {
		return nil, err
	}
	return &Multiqueue{Element: e}, nil
}

func NewMultiqueueWithName(name string) (*Multiqueue, error) {
	e, err := gst.NewElementWithName("multiqueue", name)
	if err != nil {
		return nil, err
	}
	return &Multiqueue{Element: e}, nil
}

// ----- Properties -----

// extra-size-buffers (uint)
//
// Amount of buffers the queues can grow if one of them is empty (0=disable) (NOT IMPLEMENTED)

func (e *Multiqueue) GetExtraSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("extra-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Multiqueue) SetExtraSizeBuffers(value uint) error {
	return e.Element.SetProperty("extra-size-buffers", value)
}

// extra-size-bytes (uint)
//
// Amount of data the queues can grow if one of them is empty (bytes, 0=disable) (NOT IMPLEMENTED)

func (e *Multiqueue) GetExtraSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("extra-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Multiqueue) SetExtraSizeBytes(value uint) error {
	return e.Element.SetProperty("extra-size-bytes", value)
}

// extra-size-time (uint64)
//
// Amount of time the queues can grow if one of them is empty (in ns, 0=disable) (NOT IMPLEMENTED)

func (e *Multiqueue) GetExtraSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("extra-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Multiqueue) SetExtraSizeTime(value uint64) error {
	return e.Element.SetProperty("extra-size-time", value)
}

// high-percent (int)
//
// High threshold percent for buffering to finish.

func (e *Multiqueue) GetHighPercent() (int, error) {
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

func (e *Multiqueue) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

// high-watermark (float64)
//
// High threshold watermark for buffering to finish.

func (e *Multiqueue) GetHighWatermark() (float64, error) {
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

func (e *Multiqueue) SetHighWatermark(value float64) error {
	return e.Element.SetProperty("high-watermark", value)
}

// low-percent (int)
//
// Low threshold percent for buffering to start.

func (e *Multiqueue) GetLowPercent() (int, error) {
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

func (e *Multiqueue) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

// low-watermark (float64)
//
// Low threshold watermark for buffering to start.

func (e *Multiqueue) GetLowWatermark() (float64, error) {
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

func (e *Multiqueue) SetLowWatermark(value float64) error {
	return e.Element.SetProperty("low-watermark", value)
}

// max-size-buffers (uint)
//
// Max. number of buffers in the queue (0=disable)

func (e *Multiqueue) GetMaxSizeBuffers() (uint, error) {
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

func (e *Multiqueue) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

// max-size-bytes (uint)
//
// Max. amount of data in the queue (bytes, 0=disable)

func (e *Multiqueue) GetMaxSizeBytes() (uint, error) {
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

func (e *Multiqueue) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

// max-size-time (uint64)
//
// Max. amount of data in the queue (in ns, 0=disable)

func (e *Multiqueue) GetMaxSizeTime() (uint64, error) {
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

func (e *Multiqueue) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

// min-interleave-time (uint64)
//
// Minimum extra buffering for deinterleaving (size of the queues) when use-interleave=true

func (e *Multiqueue) GetMinInterleaveTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-interleave-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Multiqueue) SetMinInterleaveTime(value uint64) error {
	return e.Element.SetProperty("min-interleave-time", value)
}

// stats (GstStructure)
//
// Various multiqueue statistics. This property returns a GstStructure
// with name "application/x-gst-multi-queue-stats" with the following fields:

// sync-by-running-time (bool)
//
// If enabled multiqueue will synchronize deactivated or not-linked streams
// to the activated and linked streams by taking the running time.
// Otherwise multiqueue will synchronize the deactivated or not-linked
// streams by keeping the order in which buffers and events arrived compared
// to active and linked streams.

func (e *Multiqueue) GetSyncByRunningTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync-by-running-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiqueue) SetSyncByRunningTime(value bool) error {
	return e.Element.SetProperty("sync-by-running-time", value)
}

// unlinked-cache-time (uint64)
//
// Extra buffering in time for unlinked streams (if 'sync-by-running-time')

func (e *Multiqueue) GetUnlinkedCacheTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("unlinked-cache-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Multiqueue) SetUnlinkedCacheTime(value uint64) error {
	return e.Element.SetProperty("unlinked-cache-time", value)
}

// use-buffering (bool)
//
// Enable the buffering option in multiqueue so that BUFFERING messages are
// emitted based on low-/high-percent thresholds.

func (e *Multiqueue) GetUseBuffering() (bool, error) {
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

func (e *Multiqueue) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

// use-interleave (bool)
//
// Adjust time limits based on input interleave

func (e *Multiqueue) GetUseInterleave() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-interleave")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multiqueue) SetUseInterleave(value bool) error {
	return e.Element.SetProperty("use-interleave", value)
}

