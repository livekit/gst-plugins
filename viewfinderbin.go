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

type Viewfinderbin struct {
	Element *gst.Element
}

func NewViewfinderbin() (*Viewfinderbin, error) {
	e, err := gst.NewElement("viewfinderbin")
	if err != nil {
		return nil, err
	}
	return &Viewfinderbin{Element: e}, nil
}

func NewViewfinderbinWithName(name string) (*Viewfinderbin, error) {
	e, err := gst.NewElementWithName("viewfinderbin", name)
	if err != nil {
		return nil, err
	}
	return &Viewfinderbin{Element: e}, nil
}

// ----- Properties -----

// disable-converters (bool)
//
// If video converters should be disabled (must be set on NULL)

func (e *Viewfinderbin) GetDisableConverters() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("disable-converters")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Viewfinderbin) SetDisableConverters(value bool) error {
	return e.Element.SetProperty("disable-converters", value)
}

// video-sink (GstElement)
//
// the video output element to use (NULL = default)

func (e *Viewfinderbin) GetVideoSink() (interface{}, error) {
	return e.Element.GetProperty("video-sink")
}

func (e *Viewfinderbin) SetVideoSink(value interface{}) error {
	return e.Element.SetProperty("video-sink", value)
}

