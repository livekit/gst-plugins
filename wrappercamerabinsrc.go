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
	"github.com/tinyzimmer/go-gst/gst"
)

type Wrappercamerabinsrc struct {
	Element *gst.Element
}

func NewWrappercamerabinsrc() (*Wrappercamerabinsrc, error) {
	e, err := gst.NewElement("wrappercamerabinsrc")
	if err != nil {
		return nil, err
	}
	return &Wrappercamerabinsrc{Element: e}, nil
}

func NewWrappercamerabinsrcWithName(name string) (*Wrappercamerabinsrc, error) {
	e, err := gst.NewElementWithName("wrappercamerabinsrc", name)
	if err != nil {
		return nil, err
	}
	return &Wrappercamerabinsrc{Element: e}, nil
}

// ----- Properties -----

// video-source (GstElement)
//
// The video source element to be used

func (e *Wrappercamerabinsrc) GetVideoSource() (interface{}, error) {
	return e.Element.GetProperty("video-source")
}

func (e *Wrappercamerabinsrc) SetVideoSource(value interface{}) error {
	return e.Element.SetProperty("video-source", value)
}

// video-source-filter (GstElement)
//
// Optional video source filter element

func (e *Wrappercamerabinsrc) GetVideoSourceFilter() (interface{}, error) {
	return e.Element.GetProperty("video-source-filter")
}

func (e *Wrappercamerabinsrc) SetVideoSourceFilter(value interface{}) error {
	return e.Element.SetProperty("video-source-filter", value)
}

