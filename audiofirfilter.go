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

type Audiofirfilter struct {
	*base.GstBaseTransform
}

func NewAudiofirfilter() (*Audiofirfilter, error) {
	e, err := gst.NewElement("audiofirfilter")
	if err != nil {
		return nil, err
	}
	return &Audiofirfilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAudiofirfilterWithName(name string) (*Audiofirfilter, error) {
	e, err := gst.NewElementWithName("audiofirfilter", name)
	if err != nil {
		return nil, err
	}
	return &Audiofirfilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// kernel (GstGvalueArray)
//
// Filter kernel for the FIR filter

func (e *Audiofirfilter) GetKernel() (interface{}, error) {
	return e.Element.GetProperty("kernel")
}

func (e *Audiofirfilter) SetKernel(value interface{}) error {
	return e.Element.SetProperty("kernel", value)
}

// latency (uint64)
//
// Filter latency in samples

func (e *Audiofirfilter) GetLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Audiofirfilter) SetLatency(value uint64) error {
	return e.Element.SetProperty("latency", value)
}

