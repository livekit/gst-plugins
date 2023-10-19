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

type Rtph264Depay struct {
	Element *gst.Element
}

func NewRtph264Depay() (*Rtph264Depay, error) {
	e, err := gst.NewElement("rtph264depay")
	if err != nil {
		return nil, err
	}
	return &Rtph264Depay{Element: e}, nil
}

func NewRtph264DepayWithName(name string) (*Rtph264Depay, error) {
	e, err := gst.NewElementWithName("rtph264depay", name)
	if err != nil {
		return nil, err
	}
	return &Rtph264Depay{Element: e}, nil
}

// ----- Properties -----

// request-keyframe (bool)
//
// Request new keyframe when packet loss is detected

func (e *Rtph264Depay) GetRequestKeyframe() (bool, error) {
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

func (e *Rtph264Depay) SetRequestKeyframe(value bool) error {
	return e.Element.SetProperty("request-keyframe", value)
}

// wait-for-keyframe (bool)
//
// Wait for the next keyframe after packet loss,
// meaningful only when outputting access units

func (e *Rtph264Depay) GetWaitForKeyframe() (bool, error) {
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

func (e *Rtph264Depay) SetWaitForKeyframe(value bool) error {
	return e.Element.SetProperty("wait-for-keyframe", value)
}

