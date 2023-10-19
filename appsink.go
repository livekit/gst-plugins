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

type Appsink struct {
	*base.GstBaseSink
}

func NewAppsink() (*Appsink, error) {
	e, err := gst.NewElement("appsink")
	if err != nil {
		return nil, err
	}
	return &Appsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAppsinkWithName(name string) (*Appsink, error) {
	e, err := gst.NewElementWithName("appsink", name)
	if err != nil {
		return nil, err
	}
	return &Appsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// buffer-list (bool)
//
// Use buffer lists

func (e *Appsink) GetBufferList() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("buffer-list")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Appsink) SetBufferList(value bool) error {
	return e.Element.SetProperty("buffer-list", value)
}

// caps (GstCaps)
//
// The allowed caps for the sink pad

func (e *Appsink) GetCaps() (*gst.Caps, error) {
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

func (e *Appsink) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// drop (bool)
//
// Drop old buffers when the buffer queue is filled

func (e *Appsink) GetDrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Appsink) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

// emit-signals (bool)
//
// Emit new-preroll and new-sample signals

func (e *Appsink) GetEmitSignals() (bool, error) {
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

func (e *Appsink) SetEmitSignals(value bool) error {
	return e.Element.SetProperty("emit-signals", value)
}

// eos (bool)
//
// Check if the sink is EOS or not started

func (e *Appsink) GetEos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// max-buffers (uint)
//
// Maximum amount of buffers in the queue (0 = unlimited).

func (e *Appsink) GetMaxBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Appsink) SetMaxBuffers(value uint) error {
	return e.Element.SetProperty("max-buffers", value)
}

// max-bytes (uint64)
//
// Maximum amount of bytes in the queue (0 = unlimited)

func (e *Appsink) GetMaxBytes() (uint64, error) {
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

func (e *Appsink) SetMaxBytes(value uint64) error {
	return e.Element.SetProperty("max-bytes", value)
}

// max-time (uint64)
//
// Maximum total duration of data in the queue (0 = unlimited)

func (e *Appsink) GetMaxTime() (uint64, error) {
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

func (e *Appsink) SetMaxTime(value uint64) error {
	return e.Element.SetProperty("max-time", value)
}

// wait-on-eos (bool)
//
// Wait for all buffers to be processed after receiving an EOS.

