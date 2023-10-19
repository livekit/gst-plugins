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

type Grabcut struct {
	*base.GstBaseTransform
}

func NewGrabcut() (*Grabcut, error) {
	e, err := gst.NewElement("grabcut")
	if err != nil {
		return nil, err
	}
	return &Grabcut{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGrabcutWithName(name string) (*Grabcut, error) {
	e, err := gst.NewElementWithName("grabcut", name)
	if err != nil {
		return nil, err
	}
	return &Grabcut{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// scale (float32)
//
// Grow factor for the face bounding box, if present

func (e *Grabcut) GetScale() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Grabcut) SetScale(value float32) error {
	return e.Element.SetProperty("scale", value)
}

// test-mode (bool)
//
// If true, the output RGB is overwritten with the segmented foreground. Alpha channel same as normal case

func (e *Grabcut) GetTestMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("test-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Grabcut) SetTestMode(value bool) error {
	return e.Element.SetProperty("test-mode", value)
}

