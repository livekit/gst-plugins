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

type Cameraundistort struct {
	*base.GstBaseTransform
}

func NewCameraundistort() (*Cameraundistort, error) {
	e, err := gst.NewElement("cameraundistort")
	if err != nil {
		return nil, err
	}
	return &Cameraundistort{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCameraundistortWithName(name string) (*Cameraundistort, error) {
	e, err := gst.NewElementWithName("cameraundistort", name)
	if err != nil {
		return nil, err
	}
	return &Cameraundistort{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// alpha (float32)
//
// Show all pixels (1), only valid ones (0) or something in between

func (e *Cameraundistort) GetAlpha() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("alpha")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Cameraundistort) SetAlpha(value float32) error {
	return e.Element.SetProperty("alpha", value)
}

// settings (string)
//
// Camera correction parameters (opaque string of serialized OpenCV objects)

func (e *Cameraundistort) GetSettings() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("settings")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Cameraundistort) SetSettings(value string) error {
	return e.Element.SetProperty("settings", value)
}

// undistort (bool)
//
// Apply camera corrections

func (e *Cameraundistort) GetUndistort() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("undistort")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Cameraundistort) SetUndistort(value bool) error {
	return e.Element.SetProperty("undistort", value)
}

