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

type Rtpvp9Pay struct {
	Element *gst.Element
}

func NewRtpvp9Pay() (*Rtpvp9Pay, error) {
	e, err := gst.NewElement("rtpvp9pay")
	if err != nil {
		return nil, err
	}
	return &Rtpvp9Pay{Element: e}, nil
}

func NewRtpvp9PayWithName(name string) (*Rtpvp9Pay, error) {
	e, err := gst.NewElementWithName("rtpvp9pay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpvp9Pay{Element: e}, nil
}

// ----- Properties -----

// picture-id (int)
//
// Currently used picture-id

func (e *Rtpvp9Pay) GetPictureId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("picture-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// picture-id-mode (GstVp9RtppayMode)
//
// The picture ID mode for payloading

func (e *Rtpvp9Pay) GetPictureIdMode() (GstVp9RtppayMode, error) {
	var value GstVp9RtppayMode
	var ok bool
	v, err := e.Element.GetProperty("picture-id-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVp9RtppayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVp9RtppayMode")
	}
	return value, nil
}

func (e *Rtpvp9Pay) SetPictureIdMode(value GstVp9RtppayMode) error {
	e.Element.SetArg("picture-id-mode", string(value))
	return nil
}

// picture-id-offset (int)
//
// Offset to add to the initial picture-id (-1 = random)

func (e *Rtpvp9Pay) GetPictureIdOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("picture-id-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpvp9Pay) SetPictureIdOffset(value int) error {
	return e.Element.SetProperty("picture-id-offset", value)
}

// ----- Constants -----

type GstVp9RtppayMode string

const (
	GstVp9RtppayModeNone GstVp9RtppayMode = "none" // none (0) – No Picture ID
	GstVp9RtppayMode7Bit GstVp9RtppayMode = "7-bit" // 7-bit (1) – 7-bit Picture ID
	GstVp9RtppayMode15Bit GstVp9RtppayMode = "15-bit" // 15-bit (2) – 15-bit Picture ID
)

