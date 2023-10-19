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

type Rtphdrexttwcc struct {
	Element *gst.Element
}

func NewRtphdrexttwcc() (*Rtphdrexttwcc, error) {
	e, err := gst.NewElement("rtphdrexttwcc")
	if err != nil {
		return nil, err
	}
	return &Rtphdrexttwcc{Element: e}, nil
}

func NewRtphdrexttwccWithName(name string) (*Rtphdrexttwcc, error) {
	e, err := gst.NewElementWithName("rtphdrexttwcc", name)
	if err != nil {
		return nil, err
	}
	return &Rtphdrexttwcc{Element: e}, nil
}

// ----- Properties -----

// n-streams (uint)
//
// The number of independant RTP streams that are being used for the transport
// wide counter for TWCC.  If set to 1 (the default), then any existing
// transport wide counter is kept.

func (e *Rtphdrexttwcc) GetNStreams() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("n-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtphdrexttwcc) SetNStreams(value uint) error {
	return e.Element.SetProperty("n-streams", value)
}

