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

type Intervideosrc struct {
	*base.GstBaseSrc
}

func NewIntervideosrc() (*Intervideosrc, error) {
	e, err := gst.NewElement("intervideosrc")
	if err != nil {
		return nil, err
	}
	return &Intervideosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewIntervideosrcWithName(name string) (*Intervideosrc, error) {
	e, err := gst.NewElementWithName("intervideosrc", name)
	if err != nil {
		return nil, err
	}
	return &Intervideosrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// channel (string)
//
// Channel name to match inter src and sink elements

func (e *Intervideosrc) GetChannel() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("channel")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Intervideosrc) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

// timeout (uint64)
//
// Timeout after which to start outputting black frames

func (e *Intervideosrc) GetTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Intervideosrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

