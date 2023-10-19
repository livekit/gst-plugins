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

	"github.com/livekit/gstplugins/base"
)

type Giostreamsink struct {
	*base.GstBaseSink
}

func NewGiostreamsink() (*Giostreamsink, error) {
	e, err := gst.NewElement("giostreamsink")
	if err != nil {
		return nil, err
	}
	return &Giostreamsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGiostreamsinkWithName(name string) (*Giostreamsink, error) {
	e, err := gst.NewElementWithName("giostreamsink", name)
	if err != nil {
		return nil, err
	}
	return &Giostreamsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// stream (GstGoutputStream)
//
// Stream to write to

func (e *Giostreamsink) GetStream() (interface{}, error) {
	return e.Element.GetProperty("stream")
}

func (e *Giostreamsink) SetStream(value interface{}) error {
	return e.Element.SetProperty("stream", value)
}

