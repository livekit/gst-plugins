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

type Rtpssrcdemux struct {
	Element *gst.Element
}

func NewRtpssrcdemux() (*Rtpssrcdemux, error) {
	e, err := gst.NewElement("rtpssrcdemux")
	if err != nil {
		return nil, err
	}
	return &Rtpssrcdemux{Element: e}, nil
}

func NewRtpssrcdemuxWithName(name string) (*Rtpssrcdemux, error) {
	e, err := gst.NewElementWithName("rtpssrcdemux", name)
	if err != nil {
		return nil, err
	}
	return &Rtpssrcdemux{Element: e}, nil
}

// ----- Properties -----

// max-streams (uint)
//
// The maximum number of streams allowed

func (e *Rtpssrcdemux) GetMaxStreams() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpssrcdemux) SetMaxStreams(value uint) error {
	return e.Element.SetProperty("max-streams", value)
}

