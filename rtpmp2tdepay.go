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

type Rtpmp2Tdepay struct {
	Element *gst.Element
}

func NewRtpmp2Tdepay() (*Rtpmp2Tdepay, error) {
	e, err := gst.NewElement("rtpmp2tdepay")
	if err != nil {
		return nil, err
	}
	return &Rtpmp2Tdepay{Element: e}, nil
}

func NewRtpmp2TdepayWithName(name string) (*Rtpmp2Tdepay, error) {
	e, err := gst.NewElementWithName("rtpmp2tdepay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpmp2Tdepay{Element: e}, nil
}

// ----- Properties -----

// skip-first-bytes (uint)
//
// The amount of bytes that need to be skipped at the beginning of the payload

func (e *Rtpmp2Tdepay) GetSkipFirstBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("skip-first-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpmp2Tdepay) SetSkipFirstBytes(value uint) error {
	return e.Element.SetProperty("skip-first-bytes", value)
}

