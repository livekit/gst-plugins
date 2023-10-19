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

type Mpegtsmux struct {
	*base.GstAggregator
}

func NewMpegtsmux() (*Mpegtsmux, error) {
	e, err := gst.NewElement("mpegtsmux")
	if err != nil {
		return nil, err
	}
	return &Mpegtsmux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewMpegtsmuxWithName(name string) (*Mpegtsmux, error) {
	e, err := gst.NewElementWithName("mpegtsmux", name)
	if err != nil {
		return nil, err
	}
	return &Mpegtsmux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// m2ts-mode (bool)
//
// Set to TRUE to output Blu-Ray disc format with 192 byte packets. FALSE for standard TS format with 188 byte packets.

func (e *Mpegtsmux) GetM2TsMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m2ts-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mpegtsmux) SetM2TsMode(value bool) error {
	return e.Element.SetProperty("m2ts-mode", value)
}

