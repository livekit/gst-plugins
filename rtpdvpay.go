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

type Rtpdvpay struct {
	Element *gst.Element
}

func NewRtpdvpay() (*Rtpdvpay, error) {
	e, err := gst.NewElement("rtpdvpay")
	if err != nil {
		return nil, err
	}
	return &Rtpdvpay{Element: e}, nil
}

func NewRtpdvpayWithName(name string) (*Rtpdvpay, error) {
	e, err := gst.NewElementWithName("rtpdvpay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpdvpay{Element: e}, nil
}

// ----- Properties -----

// mode (GstDvpayMode)
//
// The payload mode of payloading

func (e *Rtpdvpay) GetMode() (GstDvpayMode, error) {
	var value GstDvpayMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvpayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvpayMode")
	}
	return value, nil
}

func (e *Rtpdvpay) SetMode(value GstDvpayMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ----- Constants -----

type GstDvpayMode string

const (
	GstDvpayModeVideo GstDvpayMode = "video" // video (0) – Video only
	GstDvpayModeBundled GstDvpayMode = "bundled" // bundled (1) – Video and Audio bundled
	GstDvpayModeAudio GstDvpayMode = "audio" // audio (2) – Audio only
)

