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

type Rtpfunnel struct {
	Element *gst.Element
}

func NewRtpfunnel() (*Rtpfunnel, error) {
	e, err := gst.NewElement("rtpfunnel")
	if err != nil {
		return nil, err
	}
	return &Rtpfunnel{Element: e}, nil
}

func NewRtpfunnelWithName(name string) (*Rtpfunnel, error) {
	e, err := gst.NewElementWithName("rtpfunnel", name)
	if err != nil {
		return nil, err
	}
	return &Rtpfunnel{Element: e}, nil
}

// ----- Properties -----

// common-ts-offset (int)
//
// Use the same RTP timestamp offset for all sinkpads (-1 = disable)

func (e *Rtpfunnel) GetCommonTsOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("common-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpfunnel) SetCommonTsOffset(value int) error {
	return e.Element.SetProperty("common-ts-offset", value)
}

