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

type Vulkanshaderspv struct {
	*base.GstBaseTransform
}

func NewVulkanshaderspv() (*Vulkanshaderspv, error) {
	e, err := gst.NewElement("vulkanshaderspv")
	if err != nil {
		return nil, err
	}
	return &Vulkanshaderspv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewVulkanshaderspvWithName(name string) (*Vulkanshaderspv, error) {
	e, err := gst.NewElementWithName("vulkanshaderspv", name)
	if err != nil {
		return nil, err
	}
	return &Vulkanshaderspv{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// fragment (GstGbytes)
//
// SPIRV fragment binary

func (e *Vulkanshaderspv) GetFragment() (interface{}, error) {
	return e.Element.GetProperty("fragment")
}

func (e *Vulkanshaderspv) SetFragment(value interface{}) error {
	return e.Element.SetProperty("fragment", value)
}

// fragment-location (string)
//
// SPIRV fragment source

func (e *Vulkanshaderspv) GetFragmentLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fragment-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Vulkanshaderspv) SetFragmentLocation(value string) error {
	return e.Element.SetProperty("fragment-location", value)
}

// vertex (GstGbytes)
//
// SPIRV vertex binary

func (e *Vulkanshaderspv) GetVertex() (interface{}, error) {
	return e.Element.GetProperty("vertex")
}

func (e *Vulkanshaderspv) SetVertex(value interface{}) error {
	return e.Element.SetProperty("vertex", value)
}

// vertex-location (string)
//
// SPIRV vertex source

func (e *Vulkanshaderspv) GetVertexLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("vertex-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Vulkanshaderspv) SetVertexLocation(value string) error {
	return e.Element.SetProperty("vertex-location", value)
}

