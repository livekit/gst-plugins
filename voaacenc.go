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

type Voaacenc struct {
	Element *gst.Element
}

func NewVoaacenc() (*Voaacenc, error) {
	e, err := gst.NewElement("voaacenc")
	if err != nil {
		return nil, err
	}
	return &Voaacenc{Element: e}, nil
}

func NewVoaacencWithName(name string) (*Voaacenc, error) {
	e, err := gst.NewElementWithName("voaacenc", name)
	if err != nil {
		return nil, err
	}
	return &Voaacenc{Element: e}, nil
}

// ----- Properties -----

// bitrate (int)
//
// Target Audio Bitrate (bits per second)

func (e *Voaacenc) GetBitrate() (int, error) {
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

func (e *Voaacenc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// hard-resync (bool)
//
// Perform clipping and sample flushing upon discontinuity

func (e *Voaacenc) GetHardResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hard-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Voaacenc) SetHardResync(value bool) error {
	return e.Element.SetProperty("hard-resync", value)
}

// mark-granule (bool)
//
// Apply granule semantics to buffer metadata (implies perfect-timestamp)

func (e *Voaacenc) GetMarkGranule() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mark-granule")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// perfect-timestamp (bool)
//
// Favour perfect timestamps over tracking upstream timestamps

func (e *Voaacenc) GetPerfectTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("perfect-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Voaacenc) SetPerfectTimestamp(value bool) error {
	return e.Element.SetProperty("perfect-timestamp", value)
}

// tolerance (int64)
//
// Consider discontinuity if timestamp jitter/imperfection exceeds tolerance (ns)

func (e *Voaacenc) GetTolerance() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Voaacenc) SetTolerance(value int64) error {
	return e.Element.SetProperty("tolerance", value)
}

