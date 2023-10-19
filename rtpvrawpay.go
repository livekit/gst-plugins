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

type Rtpvrawpay struct {
	Element *gst.Element
}

func NewRtpvrawpay() (*Rtpvrawpay, error) {
	e, err := gst.NewElement("rtpvrawpay")
	if err != nil {
		return nil, err
	}
	return &Rtpvrawpay{Element: e}, nil
}

func NewRtpvrawpayWithName(name string) (*Rtpvrawpay, error) {
	e, err := gst.NewElementWithName("rtpvrawpay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpvrawpay{Element: e}, nil
}

// ----- Properties -----

// chunks-per-frame (int)
//
// Split and send out each frame in multiple chunks to reduce overhead

func (e *Rtpvrawpay) GetChunksPerFrame() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("chunks-per-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpvrawpay) SetChunksPerFrame(value int) error {
	return e.Element.SetProperty("chunks-per-frame", value)
}

