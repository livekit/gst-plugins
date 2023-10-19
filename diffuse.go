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

type Diffuse struct {
	*base.GstBaseTransform
}

func NewDiffuse() (*Diffuse, error) {
	e, err := gst.NewElement("diffuse")
	if err != nil {
		return nil, err
	}
	return &Diffuse{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDiffuseWithName(name string) (*Diffuse, error) {
	e, err := gst.NewElementWithName("diffuse", name)
	if err != nil {
		return nil, err
	}
	return &Diffuse{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// scale (float64)
//
// Scale of the texture

func (e *Diffuse) GetScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Diffuse) SetScale(value float64) error {
	return e.Element.SetProperty("scale", value)
}

