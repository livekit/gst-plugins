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

type Flitetestsrc struct {
	*base.GstBaseSrc
}

func NewFlitetestsrc() (*Flitetestsrc, error) {
	e, err := gst.NewElement("flitetestsrc")
	if err != nil {
		return nil, err
	}
	return &Flitetestsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewFlitetestsrcWithName(name string) (*Flitetestsrc, error) {
	e, err := gst.NewElementWithName("flitetestsrc", name)
	if err != nil {
		return nil, err
	}
	return &Flitetestsrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// samplesperbuffer (int)
//
// Number of samples in each outgoing buffer

func (e *Flitetestsrc) GetSamplesperbuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("samplesperbuffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Flitetestsrc) SetSamplesperbuffer(value int) error {
	return e.Element.SetProperty("samplesperbuffer", value)
}

