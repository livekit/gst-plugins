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

	"github.com/livekit/gstplugins/base"
)

type Mirror struct {
	*base.GstBaseTransform
}

func NewMirror() (*Mirror, error) {
	e, err := gst.NewElement("mirror")
	if err != nil {
		return nil, err
	}
	return &Mirror{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewMirrorWithName(name string) (*Mirror, error) {
	e, err := gst.NewElementWithName("mirror", name)
	if err != nil {
		return nil, err
	}
	return &Mirror{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// mode (GstMirrorMode)
//
// How to split the video frame and which side reflect

func (e *Mirror) GetMode() (GstMirrorMode, error) {
	var value GstMirrorMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMirrorMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMirrorMode")
	}
	return value, nil
}

func (e *Mirror) SetMode(value GstMirrorMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// ----- Constants -----

type GstMirrorMode string

const (
	GstMirrorModeLeft GstMirrorMode = "left" // left (0) – Split horizontally and reflect left into right
	GstMirrorModeRight GstMirrorMode = "right" // right (1) – Split horizontally and reflect right into left
	GstMirrorModeTop GstMirrorMode = "top" // top (2) – Split vertically and reflect top into bottom
	GstMirrorModeBottom GstMirrorMode = "bottom" // bottom (3) – Split vertically and reflect bottom into top
)

