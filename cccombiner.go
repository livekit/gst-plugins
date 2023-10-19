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

type Cccombiner struct {
	*base.GstAggregator
}

func NewCccombiner() (*Cccombiner, error) {
	e, err := gst.NewElement("cccombiner")
	if err != nil {
		return nil, err
	}
	return &Cccombiner{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewCccombinerWithName(name string) (*Cccombiner, error) {
	e, err := gst.NewElementWithName("cccombiner", name)
	if err != nil {
		return nil, err
	}
	return &Cccombiner{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// max-scheduled (uint)
//
// Controls the number of scheduled buffers after which the element
// will start dropping old buffers from its internal queues. See
// schedule.

func (e *Cccombiner) GetMaxScheduled() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-scheduled")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Cccombiner) SetMaxScheduled(value uint) error {
	return e.Element.SetProperty("max-scheduled", value)
}

// output-padding (bool)
//
// When schedule is TRUE, this property controls
// whether the output closed caption meta stream will be padded.

func (e *Cccombiner) GetOutputPadding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cccombiner) SetOutputPadding(value bool) error {
	return e.Element.SetProperty("output-padding", value)
}

// schedule (bool)
//
// Controls whether caption buffers should be smoothly scheduled
// in order to have exactly one per output video buffer.

