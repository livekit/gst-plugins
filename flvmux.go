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

type Flvmux struct {
	*base.GstAggregator
}

func NewFlvmux() (*Flvmux, error) {
	e, err := gst.NewElement("flvmux")
	if err != nil {
		return nil, err
	}
	return &Flvmux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewFlvmuxWithName(name string) (*Flvmux, error) {
	e, err := gst.NewElementWithName("flvmux", name)
	if err != nil {
		return nil, err
	}
	return &Flvmux{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// encoder (string)
//
// The value of encoder in the meta packet.

func (e *Flvmux) GetEncoder() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("encoder")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Flvmux) SetEncoder(value string) error {
	return e.Element.SetProperty("encoder", value)
}

// enforce-increasing-timestamps (bool)
//
// If set to true, flvmux will modify buffers timestamps to ensure they are always
// strictly increasing, inside one stream and also between the audio and video streams.

func (e *Flvmux) GetEnforceIncreasingTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enforce-increasing-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Flvmux) SetEnforceIncreasingTimestamps(value bool) error {
	return e.Element.SetProperty("enforce-increasing-timestamps", value)
}

// metadatacreator (string)
//
// The value of metadatacreator in the meta packet.

func (e *Flvmux) GetMetadatacreator() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("metadatacreator")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Flvmux) SetMetadatacreator(value string) error {
	return e.Element.SetProperty("metadatacreator", value)
}

// skip-backwards-streams (bool)
//
// If set to true, streams that go backwards related to the other stream will have buffers dropped until they reach the correct timestamp

func (e *Flvmux) GetSkipBackwardsStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-backwards-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Flvmux) SetSkipBackwardsStreams(value bool) error {
	return e.Element.SetProperty("skip-backwards-streams", value)
}

// streamable (bool)
//
// If True, the output will be streaming friendly. (ie without indexes and
// duration)

func (e *Flvmux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Flvmux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

