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

type Liveadder struct {
	*base.GstAggregator
}

func NewLiveadder() (*Liveadder, error) {
	e, err := gst.NewElement("liveadder")
	if err != nil {
		return nil, err
	}
	return &Liveadder{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewLiveadderWithName(name string) (*Liveadder, error) {
	e, err := gst.NewElementWithName("liveadder", name)
	if err != nil {
		return nil, err
	}
	return &Liveadder{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// latency (uint)
//
// Additional latency in live mode to allow upstream to take longer to produce buffers for the current position (in milliseconds)

func (e *Liveadder) GetLatency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Liveadder) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

