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

type Parsebin struct {
	Element *gst.Element
}

func NewParsebin() (*Parsebin, error) {
	e, err := gst.NewElement("parsebin")
	if err != nil {
		return nil, err
	}
	return &Parsebin{Element: e}, nil
}

func NewParsebinWithName(name string) (*Parsebin, error) {
	e, err := gst.NewElementWithName("parsebin", name)
	if err != nil {
		return nil, err
	}
	return &Parsebin{Element: e}, nil
}

// ----- Properties -----

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Parsebin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Parsebin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// expose-all-streams (bool)
//
// Expose all streams, including those of unknown type or that don't match the 'caps' property

func (e *Parsebin) GetExposeAllStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-all-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Parsebin) SetExposeAllStreams(value bool) error {
	return e.Element.SetProperty("expose-all-streams", value)
}

// sink-caps (GstCaps)
//
// The caps of the input data. (NULL = use typefind element)

func (e *Parsebin) GetSinkCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("sink-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Parsebin) SetSinkCaps(value *gst.Caps) error {
	return e.Element.SetProperty("sink-caps", value)
}

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Parsebin) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Parsebin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

