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

type Rganalysis struct {
	*base.GstBaseTransform
}

func NewRganalysis() (*Rganalysis, error) {
	e, err := gst.NewElement("rganalysis")
	if err != nil {
		return nil, err
	}
	return &Rganalysis{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewRganalysisWithName(name string) (*Rganalysis, error) {
	e, err := gst.NewElementWithName("rganalysis", name)
	if err != nil {
		return nil, err
	}
	return &Rganalysis{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// forced (bool)
//
// Whether to analyze streams even when ReplayGain tags exist.

// message (bool)
//
// Post statics messages

func (e *Rganalysis) GetMessage() (bool, error) {
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

func (e *Rganalysis) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

// num-tracks (int)
//
// Number of remaining album tracks.

// reference-level (float64)
//
// Reference level [dB].

