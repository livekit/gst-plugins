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

type Videomixer struct {
	Element *gst.Element
}

func NewVideomixer() (*Videomixer, error) {
	e, err := gst.NewElement("videomixer")
	if err != nil {
		return nil, err
	}
	return &Videomixer{Element: e}, nil
}

func NewVideomixerWithName(name string) (*Videomixer, error) {
	e, err := gst.NewElementWithName("videomixer", name)
	if err != nil {
		return nil, err
	}
	return &Videomixer{Element: e}, nil
}

// ----- Properties -----

// background (GstVideoMixer2Background)
//
// Background type

func (e *Videomixer) GetBackground() (GstVideoMixer2Background, error) {
	var value GstVideoMixer2Background
	var ok bool
	v, err := e.Element.GetProperty("background")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVideoMixer2Background)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVideoMixer2Background")
	}
	return value, nil
}

func (e *Videomixer) SetBackground(value GstVideoMixer2Background) error {
	e.Element.SetArg("background", string(value))
	return nil
}

// ----- Constants -----

type GstVideoMixer2Background string

const (
	GstVideoMixer2BackgroundChecker GstVideoMixer2Background = "checker" // checker (0) – Checker pattern
	GstVideoMixer2BackgroundBlack GstVideoMixer2Background = "black" // black (1) – Black
	GstVideoMixer2BackgroundWhite GstVideoMixer2Background = "white" // white (2) – White
	GstVideoMixer2BackgroundTransparent GstVideoMixer2Background = "transparent" // transparent (3) – Transparent Background to enable further mixing
)

