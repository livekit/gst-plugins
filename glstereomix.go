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

type Glstereomix struct {
	*base.GstAggregator
}

func NewGlstereomix() (*Glstereomix, error) {
	e, err := gst.NewElement("glstereomix")
	if err != nil {
		return nil, err
	}
	return &Glstereomix{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewGlstereomixWithName(name string) (*Glstereomix, error) {
	e, err := gst.NewElementWithName("glstereomix", name)
	if err != nil {
		return nil, err
	}
	return &Glstereomix{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// downmix-mode (GstGLStereoDownmix)
//
// Output anaglyph type to generate when downmixing to mono

func (e *Glstereomix) GetDownmixMode() (interface{}, error) {
	return e.Element.GetProperty("downmix-mode")
}

func (e *Glstereomix) SetDownmixMode(value interface{}) error {
	return e.Element.SetProperty("downmix-mode", value)
}

