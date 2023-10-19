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

type Rtpjpegpay struct {
	Element *gst.Element
}

func NewRtpjpegpay() (*Rtpjpegpay, error) {
	e, err := gst.NewElement("rtpjpegpay")
	if err != nil {
		return nil, err
	}
	return &Rtpjpegpay{Element: e}, nil
}

func NewRtpjpegpayWithName(name string) (*Rtpjpegpay, error) {
	e, err := gst.NewElementWithName("rtpjpegpay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpjpegpay{Element: e}, nil
}

// ----- Properties -----

// quality (int)
//
// Quality factor on JPEG data (unused)

func (e *Rtpjpegpay) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpjpegpay) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

// type (int)
//
// Default JPEG Type, overwritten by SOF when present

func (e *Rtpjpegpay) GetType() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpjpegpay) SetType(value int) error {
	return e.Element.SetProperty("type", value)
}

