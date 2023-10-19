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

type Rtpvp8Depay struct {
	Element *gst.Element
}

func NewRtpvp8Depay() (*Rtpvp8Depay, error) {
	e, err := gst.NewElement("rtpvp8depay")
	if err != nil {
		return nil, err
	}
	return &Rtpvp8Depay{Element: e}, nil
}

func NewRtpvp8DepayWithName(name string) (*Rtpvp8Depay, error) {
	e, err := gst.NewElementWithName("rtpvp8depay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpvp8Depay{Element: e}, nil
}

// ----- Properties -----

// request-keyframe (bool)
//
// Request new keyframe when packet loss is detected

func (e *Rtpvp8Depay) GetRequestKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("request-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpvp8Depay) SetRequestKeyframe(value bool) error {
	return e.Element.SetProperty("request-keyframe", value)
}

// wait-for-keyframe (bool)
//
// Wait for the next keyframe after packet loss

func (e *Rtpvp8Depay) GetWaitForKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpvp8Depay) SetWaitForKeyframe(value bool) error {
	return e.Element.SetProperty("wait-for-keyframe", value)
}

