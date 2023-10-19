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

type Vorbisenc struct {
	Element *gst.Element
}

func NewVorbisenc() (*Vorbisenc, error) {
	e, err := gst.NewElement("vorbisenc")
	if err != nil {
		return nil, err
	}
	return &Vorbisenc{Element: e}, nil
}

func NewVorbisencWithName(name string) (*Vorbisenc, error) {
	e, err := gst.NewElementWithName("vorbisenc", name)
	if err != nil {
		return nil, err
	}
	return &Vorbisenc{Element: e}, nil
}

// ----- Properties -----

// bitrate (int)
//
// Attempt to encode at a bitrate averaging this (in bps). This uses the bitrate management engine, and is not recommended for most users. Quality is a better alternative. (-1 == disabled)

func (e *Vorbisenc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Vorbisenc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// last-message (string)
//
// The last status message

func (e *Vorbisenc) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// managed (bool)
//
// Enable bitrate management engine

func (e *Vorbisenc) GetManaged() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("managed")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Vorbisenc) SetManaged(value bool) error {
	return e.Element.SetProperty("managed", value)
}

// max-bitrate (int)
//
// Specify a maximum bitrate (in bps). Useful for streaming applications. (-1 == disabled)

func (e *Vorbisenc) GetMaxBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Vorbisenc) SetMaxBitrate(value int) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// min-bitrate (int)
//
// Specify a minimum bitrate (in bps). Useful for encoding for a fixed-size channel. (-1 == disabled)

func (e *Vorbisenc) GetMinBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Vorbisenc) SetMinBitrate(value int) error {
	return e.Element.SetProperty("min-bitrate", value)
}

// quality (float32)
//
// Specify quality instead of specifying a particular bitrate.

func (e *Vorbisenc) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Vorbisenc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

