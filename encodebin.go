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

type Encodebin struct {
	Element *gst.Element
}

func NewEncodebin() (*Encodebin, error) {
	e, err := gst.NewElement("encodebin")
	if err != nil {
		return nil, err
	}
	return &Encodebin{Element: e}, nil
}

func NewEncodebinWithName(name string) (*Encodebin, error) {
	e, err := gst.NewElementWithName("encodebin", name)
	if err != nil {
		return nil, err
	}
	return &Encodebin{Element: e}, nil
}

// ----- Properties -----

// audio-jitter-tolerance (uint64)
//
// Amount of timestamp jitter/imperfection to allow on audio streams before inserting/dropping samples (ns)

func (e *Encodebin) GetAudioJitterTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("audio-jitter-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Encodebin) SetAudioJitterTolerance(value uint64) error {
	return e.Element.SetProperty("audio-jitter-tolerance", value)
}

// avoid-reencoding (bool)
//
// Whether to re-encode portions of compatible video streams that lay on segment boundaries

func (e *Encodebin) GetAvoidReencoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("avoid-reencoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Encodebin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

// flags (GstEncodeBinFlags)
//
// Flags to control behaviour

func (e *Encodebin) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *Encodebin) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// profile (GstEncodingProfile)
//
// The GstEncodingProfile to use

func (e *Encodebin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

func (e *Encodebin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

// queue-buffers-max (uint)
//
// Max. number of buffers in the queue (0=disable)

func (e *Encodebin) GetQueueBuffersMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("queue-buffers-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Encodebin) SetQueueBuffersMax(value uint) error {
	return e.Element.SetProperty("queue-buffers-max", value)
}

// queue-bytes-max (uint)
//
// Max. amount of data in the queue (bytes, 0=disable)

func (e *Encodebin) GetQueueBytesMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("queue-bytes-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Encodebin) SetQueueBytesMax(value uint) error {
	return e.Element.SetProperty("queue-bytes-max", value)
}

// queue-time-max (uint64)
//
// Max. amount of data in the queue (in ns, 0=disable)

func (e *Encodebin) GetQueueTimeMax() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("queue-time-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Encodebin) SetQueueTimeMax(value uint64) error {
	return e.Element.SetProperty("queue-time-max", value)
}

