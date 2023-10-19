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

type AvtpsinkexamplePipeline struct {
	*base.GstBaseSink
}

func NewAvtpsinkexamplePipeline() (*AvtpsinkexamplePipeline, error) {
	e, err := gst.NewElement("avtpsinkExample pipeline")
	if err != nil {
		return nil, err
	}
	return &AvtpsinkexamplePipeline{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAvtpsinkexamplePipelineWithName(name string) (*AvtpsinkexamplePipeline, error) {
	e, err := gst.NewElementWithName("avtpsinkExample pipeline", name)
	if err != nil {
		return nil, err
	}
	return &AvtpsinkexamplePipeline{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// address (string)
//
// Destination MAC address from Ethernet frames

func (e *AvtpsinkexamplePipeline) GetAddress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("address")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *AvtpsinkexamplePipeline) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

// ifname (string)
//
// Network interface utilized to transmit AVTPDUs

func (e *AvtpsinkexamplePipeline) GetIfname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ifname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *AvtpsinkexamplePipeline) SetIfname(value string) error {
	return e.Element.SetProperty("ifname", value)
}

// priority (int)
//
// Priority configured into socket (SO_PRIORITY)

func (e *AvtpsinkexamplePipeline) GetPriority() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("priority")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvtpsinkexamplePipeline) SetPriority(value int) error {
	return e.Element.SetProperty("priority", value)
}

