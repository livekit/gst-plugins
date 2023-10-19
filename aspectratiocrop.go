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

type Aspectratiocrop struct {
	Element *gst.Element
}

func NewAspectratiocrop() (*Aspectratiocrop, error) {
	e, err := gst.NewElement("aspectratiocrop")
	if err != nil {
		return nil, err
	}
	return &Aspectratiocrop{Element: e}, nil
}

func NewAspectratiocropWithName(name string) (*Aspectratiocrop, error) {
	e, err := gst.NewElementWithName("aspectratiocrop", name)
	if err != nil {
		return nil, err
	}
	return &Aspectratiocrop{Element: e}, nil
}

// ----- Properties -----

// aspect-ratio (GstFraction)
//
// Target aspect-ratio of video

func (e *Aspectratiocrop) GetAspectRatio() (interface{}, error) {
	return e.Element.GetProperty("aspect-ratio")
}

func (e *Aspectratiocrop) SetAspectRatio(value interface{}) error {
	return e.Element.SetProperty("aspect-ratio", value)
}

