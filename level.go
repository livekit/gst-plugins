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

type Level struct {
	*base.GstBaseTransform
}

func NewLevel() (*Level, error) {
	e, err := gst.NewElement("level")
	if err != nil {
		return nil, err
	}
	return &Level{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLevelWithName(name string) (*Level, error) {
	e, err := gst.NewElementWithName("level", name)
	if err != nil {
		return nil, err
	}
	return &Level{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// audio-level-meta (bool)
//
// If TRUE, generate or update GstAudioLevelMeta on output buffers.

func (e *Level) GetAudioLevelMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("audio-level-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Level) SetAudioLevelMeta(value bool) error {
	return e.Element.SetProperty("audio-level-meta", value)
}

// interval (uint64)
//
// Interval of time between message posts (in nanoseconds)

func (e *Level) GetInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Level) SetInterval(value uint64) error {
	return e.Element.SetProperty("interval", value)
}

// message (bool)
//
// Post a 'level' message for each passed interval (deprecated, use the post-messages property instead)

func (e *Level) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Level) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

// peak-falloff (float64)
//
// Decay rate of decay peak after TTL (in dB/sec)

func (e *Level) GetPeakFalloff() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("peak-falloff")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Level) SetPeakFalloff(value float64) error {
	return e.Element.SetProperty("peak-falloff", value)
}

// peak-ttl (uint64)
//
// Time To Live of decay peak before it falls back (in nanoseconds)

func (e *Level) GetPeakTtl() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("peak-ttl")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Level) SetPeakTtl(value uint64) error {
	return e.Element.SetProperty("peak-ttl", value)
}

// post-messages (bool)
//
// Post messages on the bus with level information.

func (e *Level) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Level) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

