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

type Intervideosink struct {
	*base.GstBaseSink
}

func NewIntervideosink() (*Intervideosink, error) {
	e, err := gst.NewElement("intervideosink")
	if err != nil {
		return nil, err
	}
	return &Intervideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewIntervideosinkWithName(name string) (*Intervideosink, error) {
	e, err := gst.NewElementWithName("intervideosink", name)
	if err != nil {
		return nil, err
	}
	return &Intervideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// channel (string)
//
// Channel name to match inter src and sink elements

func (e *Intervideosink) GetChannel() (string, error) {
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

func (e *Intervideosink) SetChannel(value string) error {
	return e.Element.SetProperty("channel", value)
}

