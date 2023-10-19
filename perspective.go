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

type Perspective struct {
	*base.GstBaseTransform
}

func NewPerspective() (*Perspective, error) {
	e, err := gst.NewElement("perspective")
	if err != nil {
		return nil, err
	}
	return &Perspective{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewPerspectiveWithName(name string) (*Perspective, error) {
	e, err := gst.NewElementWithName("perspective", name)
	if err != nil {
		return nil, err
	}
	return &Perspective{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// matrix (GstGvalueArray)
//
// Matrix of dimension 3x3 to use in the 2D transform, passed as an array of 9 elements in row-major order

func (e *Perspective) GetMatrix() (interface{}, error) {
	return e.Element.GetProperty("matrix")
}

func (e *Perspective) SetMatrix(value interface{}) error {
	return e.Element.SetProperty("matrix", value)
}

